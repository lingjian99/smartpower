package spserver

import (
	"smartpower/pkg/utils/aesx"
	"smartpower/pkg/protocol"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

type ConnCtx struct {	
	Codec
	AesCipher  aesx.AESCrypter
	Bid        string
	RemoteAddr string
	//conn      gnet.Conn
	ctChan     chan protocol.RemoteControlResult
	ctState    uint32
	ctMux      sync.Mutex
	isBind     bool   // 是否请求注册
	Lasttime   int64  // unix
	Ciphertext []byte // 硬件 Ciphertext
}

func (c *ConnCtx) isLogin() bool {
	return c.Bid != ""
}

func (c *ConnCtx) StartRemoteCtrl() bool {
	if !atomic.CompareAndSwapUint32(&c.ctState, 0, 1) {
		return false
	}

	c.ctMux.Lock()
	defer c.ctMux.Unlock()
	c.ctChan = make(chan protocol.RemoteControlResult, 1)

	return true
}

func (c *ConnCtx) WriteCtResult(data protocol.RemoteControlResult) {
	if c.ctChan == nil {
		return
	}
	c.ctChan <- data
	close(c.ctChan)
}

func (c *ConnCtx) WaitCtResult(timeout time.Duration) (protocol.RemoteControlResult, error) {
	var (
		data protocol.RemoteControlResult
		err  error
	)

	select {
	case d := <-c.ctChan:
		data = d
	case <-time.After(timeout):
		err = errors.New("timeout")
	}
	c.StopCt()
	return data, err
}

func (c *ConnCtx) StopCt() {
	c.ctMux.Lock()
	defer c.ctMux.Unlock()
	if c.ctChan != nil {
		_, _ = <-c.ctChan
	}
	c.ctChan = nil
	c.ctState = 0
}

func (c *ConnCtx) CtRunning() bool {
	return atomic.LoadUint32(&c.ctState) == 1
}
