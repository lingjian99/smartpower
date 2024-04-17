package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type cbc struct {
	key []byte
}

const BlockSize = 16

func (e *cbc) Encrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}
	b, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	c := cipher.NewCBCEncrypter(b, iv)
	dst := make([]byte, len(src))
	c.CryptBlocks(dst, src)
	return dst, nil
}

func (e *cbc) Decrypt(src []byte) ([]byte, error) {
	if len(src)%BlockSize != 0 {
		return nil, errors.New("aesx/cbc: input not full blocks")
	}
	b, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	c := cipher.NewCBCDecrypter(b, iv)
	dst := make([]byte, len(src))
	c.CryptBlocks(dst, src)
	return dst, nil
}
