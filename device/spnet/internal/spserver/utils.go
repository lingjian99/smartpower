package spserver

import (
	"errors"
	"smartpower/pkg/protocol"
	"smartpower/pkg/utils/aesx"
)

func encryptInput(aesCipher aesx.AESCrypter, out *protocol.Input) ([]byte, error) {
	if aesCipher == nil {
		return nil, errors.New("the aesCipher is nil")
	}
	if out == nil {
		return nil, errors.New("encrypt data can not be nil")
	}

	data, err := aesCipher.Encrypt(out.Data)
	if err != nil {
		return nil, err
	}
	if out.BodyLen == 0 {
		out.BodyLen = uint16(len(data))
	}
	out.Data = data
	resp := out.Pack()
	return resp, nil
}
