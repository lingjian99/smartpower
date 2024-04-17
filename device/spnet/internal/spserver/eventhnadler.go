package spserver

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"
	"smartpower/device/spnet/types"
	"smartpower/pkg/kqueue"
	"smartpower/pkg/protocol"
	"smartpower/pkg/protocol/encodingx"
	"smartpower/pkg/utils"
	"smartpower/pkg/utils/aesx"
	"smartpower/pkg/xerr"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/panjf2000/gnet/v2/pkg/pool/goroutine"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
)

const ReadDeadline = time.Minute

var (
	//_ HandleServer = (*EventHandler)(nil)
	ErrConnNotFound = xerr.NewErrCodeMsg(13001, "设备未连接")
	ErrConnNotInit  = xerr.NewErrCodeMsg(13002, "连接未初始化")
	Zero            = struct{}{}
)

type HandleServer interface {
	RegisterDevice(in *protocol.Register) (out *protocol.Input, err error)
	HandleDecryptedData(bid string, head protocol.Head, data []byte) (out *protocol.Input, err error)
}

type IMsgHandler interface {
	//RegisterDevice(in *protocol.Register) (out *protocol.Input, err error)
	DecryptedDataHandle(bid string, head protocol.Head, data []byte) (out *protocol.Input, err error)
}

type EventHandler struct {
	gnet.BuiltinEventEngine
	g gnet.Engine

	protoAddr string
	multicore bool

	//connectedBid     sync.Map
	connectedSockets sync.Map
	connected        int32

	rsaPriKeys  map[string]*rsa.PrivateKey
	workerPool  *goroutine.Pool
	timingWheel *collection.TimingWheel

	Producer *kqueue.Producer
	//bindingWheel *collection.TimingWheel
	tickNum    int
	MaxConnect int
	//Handler    HandleServer
}

func NewLisServer(conf *types.NetConfig, producer *kqueue.Producer) *EventHandler {
	s := &EventHandler{
		protoAddr:  conf.Network + "://" + conf.ListenOn,
		multicore:  conf.Multicore,
		workerPool: goroutine.Default(),
		MaxConnect: conf.MaxConnect,
		Producer: producer,
	}


	if err := s.loadPrivateKey(conf.PriKeys); err != nil {
		panic(err)
	}
	tw, err := collection.NewTimingWheel(time.Second, 300, func(k, _ interface{}) {
		logx.Infof("conn [%s] timeout", k)
		key, ok := k.(string)
		if !ok {
			return
		}
		conn, ok := s.connectedSockets.Load(key)
		if ok {
			_ = (conn.(gnet.Conn)).Close()
		}
	})
	if err != nil {
		panic(err)
	}
	s.timingWheel = tw


	return s
}

func (s *EventHandler) loadPrivateKey(PriKeys map[string]string) error {
	//conf := s.svcCtx.Config.TcpServer
	_priKeys := make(map[string]*rsa.PrivateKey, len(PriKeys))
	for i, key := range PriKeys {
		buf, err := os.ReadFile(key)
		if err != nil {
			return err
		}
		pri, err := utils.ParseRSAPrivateKeyFromPEM(buf)
		if err != nil {
			return err
		}
		_priKeys[i] = pri
	}
	s.rsaPriKeys = _priKeys
	return nil
}


func (s *EventHandler) Stop() {
	logx.Info("stop engine...", gnet.Stop(context.Background(), s.protoAddr))
	s.timingWheel.Stop()
}

func (s *EventHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	logx.Infof("running server on %s with multi-core=%t", s.protoAddr, s.multicore)
	s.g = eng

	return
}

func (s *EventHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	remoteAddr := c.RemoteAddr().String()
	logx.Infof("%v connected.", remoteAddr)

	// 设置 conn ctx
	ctx := &ConnCtx{RemoteAddr: remoteAddr, Lasttime: time.Now().Unix()}
	c.SetContext(ctx)

	s.connectedSockets.Store(remoteAddr, c)
	atomic.AddInt32(&s.connected, 1)
	// 设置超时踢出
	err := s.timingWheel.SetTimer(remoteAddr, Zero, ReadDeadline)
	if err != nil {
		logx.Error(err)
	}

	// 请求认证登录
	//out = protocol.PackServerBidRequest()
	return
}

func (s *EventHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	logx.Infof("%v colsed. %v", c.RemoteAddr(), err)

	s.connectedSockets.Delete(c.RemoteAddr().String())
	atomic.AddInt32(&s.connected, -1)
	ctx := c.Context().(*ConnCtx)
	if ctx.Bid != "" {
		//s.connectedBid.Delete(ctx.Bid)
		logx.Debugf("==========> %s 下线", ctx.Bid)
	}
	_ = s.timingWheel.RemoveTimer(ctx.RemoteAddr)

	return
}

func (s *EventHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	ctx := c.Context().(*ConnCtx)
	//codec := ctx.Codec
	var err error
	var packets []*protocol.Input

	for {
		data, err := ctx.Decode(c)
		if err == ErrIncompletePacket {
			break
		}
		if err != nil {
			logging.Errorf("invalid packet: %v", err)
			return gnet.Close
		}
		logx.Debugf("[%s][%s]Recev: %+v", ctx.RemoteAddr, ctx.Bid, data)

		var input = new(protocol.Input)
		err = encodingx.Unmarshal(data, input)
		if err != nil {
			logx.Errorf("[%s][encodingx: %v][err:%v]", ctx.RemoteAddr, data, err)
			return gnet.Close
		}
		// 检查 Input 数据长度是否一致
		if int(input.BodyLen) != len(input.Data) {
			logx.Errorw("bodyLen err", logx.Field("data", input), logx.Field("connctx", ctx))
			return gnet.Close
		}

		//packet, _ := ctx.Encode(data)
		packets = append(packets, input)

	}

	err = s.timingWheel.SetTimer(c.RemoteAddr().String(), Zero, ReadDeadline)
	if err != nil {
		logx.Error(err)
	}

	for _, input := range packets {
		var resp []byte

		switch input.Code {
		// 硬件发送公钥 RSA 加密的 AES 密钥
		case 0xF8FF:
			resp, action = s.AuthAes(ctx, input)
		default:
			resp, action = s.HandlerEncryptedData(c, input)
		}
		if len(resp) > 0 {
			c.Write(resp)
			logx.Debugf("resp %+v, action%+v", resp, action)
		}

	}

	return
}

func (s *EventHandler) OnTick() (delay time.Duration, action gnet.Action) {
	s.tickNum++
	if s.tickNum >= 5 {
		logx.Statf("connected-count: %d", atomic.LoadInt32(&s.connected))
		s.tickNum = 0
	}
	t := time.Now().Unix()

	s.connectedSockets.Range(func(key, v any) bool {
		c := v.(gnet.Conn)
		if c == nil {
			return true
		}
		ctx := c.Context().(*ConnCtx)
		if ctx == nil {
			return true
		}
		if ctx.Bid == "" {
			if t-ctx.Lasttime < 3 { // 3s 后请求注册
				return true
			}
			// 重发注册请求
			logx.Infof("Request Bid to %s", ctx.RemoteAddr)
			if err := c.AsyncWrite(protocol.PackServerBidRequest(), nil); err != nil {
				logx.Errorf("Bid Request %+v", err)
			}
			ctx.Lasttime = t
		}

		return true
	})
	return time.Second * 2, gnet.None
}

func (s *EventHandler) Connected() int32 {
	return s.connected
}

//func (s *EventHandler) OnMessage(c gnet.Conn, bytes []byte) {
//
//}

func (s *EventHandler) AuthAes(ctx *ConnCtx, in *protocol.Input) (resp []byte, action gnet.Action) {
	priv, ok := s.rsaPriKeys[fmt.Sprintf("%d", in.Flag)]
	if !ok {
		logx.Error(fmt.Sprintf("未知的 RSA 私钥 %+v", in.Data))
	}
	ctx.Ciphertext = in.Data
	rsaOut, err := rsa.DecryptPKCS1v15(nil, priv, in.Data)

	var respData uint8

	if err != nil || len(rsaOut) < 18 {
		logx.Errorf("%+v err: %+v", ctx, err)
		respData = 0x02
	} else {
		respData = 0x01
		c, err := aesx.NewCipher(rsaOut[:16], aesx.AesType(rsaOut[16]))
		if err != nil {
			respData = 0x02
		} else {
			ctx.AesCipher = c
		}
		logx.Infof("RSA [%s][%+v]", ctx.RemoteAddr, rsaOut[:16])
	}
	resp = protocol.NewMessagePacket(0xF8FF, 0x01, 0x00, []byte{respData})
	if ctx.AesCipher == nil {
		action = gnet.Close
	}
	//_ = ctx.conn.AsyncWrite(resp, nil)
	return
}

func (s *EventHandler) HandlerEncryptedData(c gnet.Conn, in *protocol.Input) (resp []byte, action gnet.Action) {
	ctx, ok := c.Context().(*ConnCtx)
	if !ok {
		return nil, gnet.Close
	}

	if ctx.AesCipher == nil {
		logx.Errorf("[%s][CODE:%x][DATA: %v]", c.RemoteAddr().String(), in.Head.Code, in.Data)
		return
	}
	dst, err := ctx.AesCipher.Decrypt(in.Data)
	if err != nil {
		logx.Errorf("ctx.AesCipher.Decrypt err: %+v data: %+v", err, in)
		return
	}

	var out *protocol.Input
	// 将所有消息推送到 Kafka
	switch in.Code {
	case 0xf7ff:
		out, err = s.Heartbeat("")
	case 0x01ff: // 取出 BID
		deviceRegister := new(protocol.Register)
		err = encodingx.Unmarshal(dst, deviceRegister)
		if err != nil {
			logx.Errorf("encodingx.Unmarshal: err%+v", err)
		}
		// 设置 Bid
		bid := string(deviceRegister.Bid[:])
		ctx.Bid = bid

		if bid == "MJPB88888888888888" {
			return nil, gnet.Close
		}
		c.SetContext(ctx)

		logx.Infof("注册上线 %s", bid)
		// 清除其它连接
		s.removeOtherConn(ctx.RemoteAddr, bid)

		out = &protocol.Input{
			Head: protocol.NewHead(0x10, 0x01ff, 0x03, 0x10),
			Data: protocol.RegisterPacket(time.Now()),
		}
		_, err = s.DecryptedDataHandle(ctx.Bid, in.Head, dst)

	case 0xFEFF: // 硬件请求服务器同步时间
		out, err = s.Server_TimeSynchronization(dst)

	case 0xfdff:
		// 响应远程控制数据
		if in.Head.Flag == 0x02 {
			if len(dst) != 16 {
				return nil, gnet.Close
			}
			var ctResult protocol.RemoteControlResult
			_ = encodingx.Unmarshal(dst, &ctResult)
			ctx.WriteCtResult(ctResult)
		}
	default:
		out, err = s.DecryptedDataHandle(ctx.Bid, in.Head, dst)
	}
	if err != nil {
		logx.Errorf("[%s][BID=%s][CODE=%x] :%+v", c.RemoteAddr().String(), ctx.Bid, in.Head.Code, err)
	}

	if out == nil || len(out.Data) == 0 {
		return
	}
	resp, err = encryptInput(ctx.AesCipher, out)
	if err != nil {
		logx.Error(err)
		return nil, gnet.Close
	}
	return
	//_ = ctx.conn.AsyncWrite(wd, nil)
}

func (s *EventHandler) removeOtherConn(connKey string, bid string) {

	s.connectedSockets.Range(func(key, v any) bool {
		c := v.(gnet.Conn)
		if c == nil {
			return true
		}
		if key == connKey {
			return true
		}
		ctx := c.Context().(*ConnCtx)
		if ctx.Bid == bid {
			c.Close()
		}
		return true
	})
}

func (s *EventHandler) Heartbeat(bid string) (out *protocol.Input, err error) {
	return &protocol.Input{
		Head: protocol.NewHead(0x10, 0xf7ff, 0x01, 0x10),
		Data: protocol.HeartbeatBody_16,
	}, nil
}

// Server_TimeSynchronization
func (s *EventHandler) Server_TimeSynchronization(dst []byte) (out *protocol.Input, err error) {
	//var tz int
	//if err = encodingx.Unmarshal(dst, &tz); err != nil {
	//	return nil, err
	//}
	tz := int8(dst[1])
	out = protocol.NewServerTimeSynchronization(time.Now(), int(tz))
	return
}

func (s *EventHandler) Hardware_RemoteControl_resualt(dst []byte) (out *protocol.Input, err error) {
	//var tz int
	//if err = encodingx.Unmarshal(dst, &tz); err != nil {
	//	return nil, err
	//}
	tz := int8(dst[1])
	out = protocol.NewServerTimeSynchronization(time.Now(), int(tz))
	return
}

func (s *EventHandler) GetConnWithBid(bid string) (gnet.Conn, bool) {
	var cc gnet.Conn
	s.connectedSockets.Range(func(key, v any) bool {
		c := v.(gnet.Conn)
		if c == nil {
			return true
		}
		ctx := c.Context().(*ConnCtx)
		if ctx.Bid == bid {
			cc = c
			return false
		}
		return true
	})
	return cc, cc != nil
}

func (s *EventHandler) Send(bid string, in *protocol.Input) error {
	var c gnet.Conn
	c, ok := s.GetConnWithBid(bid)
	if !ok {
		logx.Errorf("can not load conn with bid: %s", bid)
		return ErrConnNotFound
	}

	ctx, ok := c.Context().(*ConnCtx)
	if !ok {
		logx.Errorf("ctx is null for bid: %s", bid)
		return ErrConnNotInit
	}
	resp, err := encryptInput(ctx.AesCipher, in)
	if err != nil {
		return errors.Wrapf(err, "encrypt data err:%+v", err)
	}
	logx.Debugf("SEND[%+v] %+v", in.Code, resp)
	return c.AsyncWrite(resp, nil)
}

func (s *EventHandler) Close(bid string) error {
	var c gnet.Conn
	c, ok := s.GetConnWithBid(bid)
	if !ok {
		return nil
	}

	return c.Close()
}

func (s *EventHandler) DecryptedDataHandle(bid string, head protocol.Head, data []byte) (out *protocol.Input, err error) {
	//topic := "mahmsg-raw"

	input := &protocol.KqMessage{
		Bid:     bid,
		Code:    head.Code,
		Flag:    head.Flag,
		DataDst: head.DataDst,
		Data:    data,
	}
	payload, err := json.Marshal(input)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	err = s.Producer.Push("", payload)
	if err != nil {
		logx.Error(err)
	}

	switch head.Code {
	case 0x0EFF:
		out = &protocol.Input{
			Head: protocol.Head{
				Code:    0x0EFF,
				Flag:    0x03,
				DataDst: 0x10,
			},
			Data: protocol.RawData16(),
		}
	}

	return
}
/* func SendMessage(c gnet.Conn, in *protocol.Input) error {
	if c == nil {
		return ErrConnNotInit
	}
	ctx, ok := c.Context().(*ConnCtx)
	if !ok {
		return ErrConnNotInit
	}
	resp, err := encryptInput(ctx.AesCipher, in)
	if err != nil {
		return errors.Wrapf(err, "encrypt data err:%+v", err)
	}
	return c.AsyncWrite(resp, nil)
} */
