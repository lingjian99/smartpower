package utils

import (
	"fmt"
	"time"
)

func JoinIP(elems [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", elems[0], elems[1], elems[2], elems[3])
}

func Date(dates [6]byte) time.Time {
	return time.Date(2000+int(dates[0]), time.Month(dates[1]), int(dates[2]),
		int(dates[3]), int(dates[4]), int(dates[5]), 0, time.Local)
}

func DateFromBytes(dates []byte) time.Time {
	if len(dates) != 6 {
		return time.Time{}
	}
	return time.Date(2000+int(dates[0]), time.Month(dates[1]), int(dates[2]),
		int(dates[3]), int(dates[4]), int(dates[5]), 0, time.Local)
}

func DateString(dates []byte) string {
	if len(dates) != 6 {
		return ""
	}
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", 2000+int(dates[0]), dates[1], dates[2], dates[3], dates[4], dates[5])
}

func TimeFormat(unix int64, layout string) string {
	return time.Unix(unix, 0).Format(layout)
}
