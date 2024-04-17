package spserver

import (
	"encoding/binary"
	"errors"
	"github.com/panjf2000/gnet/v2"
)

var ErrIncompletePacket = errors.New("incomplete packet")

const (
	bodySize = 2 //
	codeSize = 2 // 功能码
	flagSize = 1 // 功能码下 细节标志数据
	dstSize  = 1 // 数据方向 or 加密--

	headLen = bodySize + codeSize + flagSize + dstSize
)

type Codec struct{}

func (codec *Codec) Decode(c gnet.Conn) ([]byte, error) {
	bodyOffset := headLen
	buf, _ := c.Peek(bodyOffset)
	if len(buf) < bodyOffset {
		return nil, ErrIncompletePacket
	}

	bodyLen := binary.BigEndian.Uint16(buf[:bodySize])
	msgLen := bodyOffset + int(bodyLen)
	if c.InboundBuffered() < msgLen {
		return nil, ErrIncompletePacket
	}
	buf, _ = c.Peek(msgLen)
	_, _ = c.Discard(msgLen)

	return buf, nil
}
