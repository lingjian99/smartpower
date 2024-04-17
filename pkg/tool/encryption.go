package tool

import (
	"crypto/md5"
	"fmt"
)

/** 加密方式 **/

func Md5ByString(str string) string {

	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
