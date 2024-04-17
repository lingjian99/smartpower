package protocol

import (
	"encoding/binary"
	"smartpower/receiver/protocol/encodingx"
	"time"
)

const (
	CODE_DT_Server_BID_Req = 0x01FF //服务器请求硬件注册 （查询BID）
)

type Input struct {
	Head
	Data []byte
}

func (in *Input) UnPack(data []byte) error {
	if len(data) < 6 {
		return nil
	}
	in.BodyLen = binary.BigEndian.Uint16(data[:2])
	in.Code = binary.BigEndian.Uint16(data[2:4])
	in.Flag = data[4]
	in.DataDst = data[5]
	if len(data) >= 6 {
		in.Data = data[6:]
	}
	return nil
}

func (in *Input) Pack() []byte {
	in.BodyLen = uint16(len(in.Data))

	out := make([]byte, 6+len(in.Data))

	binary.BigEndian.PutUint16(out[:2], in.BodyLen)
	binary.BigEndian.PutUint16(out[2:4], in.Code)
	out[4] = in.Flag
	out[5] = in.DataDst

	if in.BodyLen > 0 {
		copy(out[6:], in.Data)
	}

	return out
}

func (in *Input) CheckDataLen() bool {
	return int(in.BodyLen) == len(in.Data)
}

func NewMessagePacket(Code uint16, Flag byte, DataDst byte, Data []byte) []byte {
	bodyLen := uint16(len(Data))
	out := make([]byte, 6+len(Data))

	binary.BigEndian.PutUint16(out[:2], bodyLen)
	binary.BigEndian.PutUint16(out[2:4], Code)
	out[4] = Flag
	out[5] = DataDst
	if bodyLen > 0 {
		copy(out[6:], Data)
	}
	return out
}

// / 服务器请求硬件注册
func PackServerBidRequest() []byte {
	b := []byte{0x00, 0x01, 0x01, 0xFF, 0x01, 0x00, 0xff}
	return b
}

func RegisterPacket(t time.Time) []byte {
	out := make([]byte, 64)
	out[0] = 0xFE

	y, m, d := t.Date()

	out[58] = byte(y - 2000)
	out[59] = byte(m)
	out[60] = byte(d)
	out[61] = byte(t.Hour())
	out[62] = byte(t.Minute())
	out[63] = byte(t.Second())

	return out
}

func ServerRedirectSite(site string) []byte {
	resp := SiteRedirect{
		Port: 3712,
		Site: [50]byte{},
		Keep: [12]byte{},
	}
	copy(resp.Site[:], site)
	out, _ := encodingx.Marshal(resp)
	return out
}
