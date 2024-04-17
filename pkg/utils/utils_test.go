package utils

import "testing"

func TestDateString(t *testing.T) {
	d := []byte{23, 3, 4, 12, 4, 3}
	s := DateString(d)
	if s != "2023-03-04 12:04:03" {
		t.Fail()
	}

	d2 := []byte{}
	if "" != DateString(d2) {
		t.Fail()
	}
}
