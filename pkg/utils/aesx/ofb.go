package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type ofb struct {
	key []byte
}

func (e ofb) Encrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}
	b, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	c := cipher.NewOFB(b, iv)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	return dst, nil
}

func (e ofb) Decrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}
	b, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	c := cipher.NewOFB(b, iv)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	return dst, nil
}
