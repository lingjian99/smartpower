package aesx

import (
	"crypto/aes"
	"github.com/pkg/errors"
)

var key = []byte("1234567890abcdef")

var iv = []byte{
	0x01, 0x01, 0x01, 0x01,
	0x01, 0x01, 0x01, 0x01,
	0x01, 0x01, 0x01, 0x01,
	0x01, 0x01, 0x01, 0x01,
}

type AesType int

var (
	AES_ECB AesType = 0
	AES_CBC AesType = 1
	AES_CFB AesType = 2
	AES_CTR AesType = 3
	AES_OFB AesType = 4
)

var UnSupportCrypterError = errors.New("unsupport Crypter Type")

type AESCrypter interface {
	Encrypt(src []byte) ([]byte, error)

	Decrypt(src []byte) ([]byte, error)
}

type Plain struct{}

func (e Plain) Encrypt(src []byte) ([]byte, error) {
	return src, nil
}

func (e Plain) Decrypt(src []byte) ([]byte, error) {
	return src, nil
}

func NewCipher(key []byte, typ AesType) (AESCrypter, error) {
	var c AESCrypter
	switch typ {
	case AES_ECB:
		c = ecb{key: key}
	case AES_CBC:
		c = &cbc{key: key}
	case AES_CFB:
		c = cfb{key: key}
	case AES_CTR:
		c = ctr{key: key}
	case AES_OFB:
		c = ofb{key: key}
	}
	if c == nil {
		return nil, UnSupportCrypterError
	}
	return c, nil
}

type ecb struct {
	key []byte
}

func (e ecb) Encrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}

	c, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(src))
	c.Encrypt(dst, src)
	return dst, nil
}

func (e ecb) Decrypt(src []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("crypto/des: input not full block")
	}
	c, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(src))
	c.Decrypt(dst, src)
	return dst, nil
}

var _ AESCrypter = (*ecb)(nil)
