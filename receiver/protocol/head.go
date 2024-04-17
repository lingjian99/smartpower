package protocol

import (
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
)

var NotEnoughData = errors.New("data not enough")

type Head struct {
	BodyLen uint16
	Code    uint16 //功能码
	Flag    byte   //
	DataDst byte   // 数据方向
}

func NewHead(len, code uint16, flag, dst uint8) Head {
	return Head{
		BodyLen: len,
		Code:    code,
		Flag:    flag,
		DataDst: dst,
	}
}

func (in *Head) Pack() []byte {

	return []byte{
		byte(in.BodyLen >> 8),
		byte(in.BodyLen),
		byte(in.Code >> 8),
		byte(in.Code),
		in.Flag,
		in.DataDst,
	}

}

func (in *Head) Unpack(data []byte) error {
	if len(data) < 6 {
		return NotEnoughData
	}
	in.BodyLen = binary.BigEndian.Uint16(data[:2])
	in.Code = binary.BigEndian.Uint16(data[2:4])
	in.Flag = data[4]
	in.DataDst = data[5]

	return nil
}

func (in *Head) String() string {
	return fmt.Sprintf("%d %#x %#x %#x", in.BodyLen, in.Code, in.Flag, in.DataDst)
}
