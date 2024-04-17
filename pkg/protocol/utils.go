package protocol

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

func GBKToUTF8Bytes(d []byte) ([]byte, error) {
	enc := simplifiedchinese.GBK.NewDecoder()
	return enc.Bytes(d)
}

func GBKToUTF8String(d []byte) string {
	enc := simplifiedchinese.GBK.NewDecoder()
	b, _ := enc.Bytes(d)
	return string(b)
}

func TrimNULL(s []byte) []byte {
	i := bytes.IndexByte(s, 0)
	if i < 0 {
		return s
	}
	return s[:i]
	//return bytes.Trim(s, "\x00")
}

//func NullTermToString(b []byte) (s string) {
//	i := bytes.IndexByte(b, '0')
//	s = string(b[0:i])
//	return
//}

func DateFromBytes(dates []byte) time.Time {
	if len(dates) != 6 {
		return time.Time{}
	}
	return time.Date(2000+int(dates[0]), time.Month(dates[1]), int(dates[2]),
		int(dates[3]), int(dates[4]), int(dates[5]), 0, time.Local)
}
