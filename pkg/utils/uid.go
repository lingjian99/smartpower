package utils

import (
	"fmt"
	"strings"
)

func Uid(b []byte) string {
	var s string
	switch len(b) {
	case 8:
		if b[6] == 0 && b[7] == 0 {
			s = fmt.Sprintf("%x", b[:6])
		} else {
			s = fmt.Sprintf("%x", b)
		}
	default:
		s = fmt.Sprintf("%x", b)
	}
	return strings.ToUpper(s)
}
