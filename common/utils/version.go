package utils

import (
	"strconv"
	"strings"
)

func Version(ver string) (major, minor, buildId int64) {
	vers := strings.Split(ver, ".")
	lenVer := len(vers)
	if lenVer > 0 {
		major, _ = strconv.ParseInt(vers[0], 10, 64)
	}
	if lenVer > 1 {
		minor, _ = strconv.ParseInt(vers[1], 10, 64)
	}
	if lenVer > 2 {
		buildId, _ = strconv.ParseInt(vers[2], 10, 64)
	}
	return
}
