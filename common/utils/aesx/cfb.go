package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type cfb struct {
	key []byte
}

func (e cfb) Encrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}
	b, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	c := cipher.NewCFBEncrypter(b, iv)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	return dst, nil
}

func (e cfb) Decrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}
	b, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	c := cipher.NewCFBDecrypter(b, iv)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	return dst, nil
}
