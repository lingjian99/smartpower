package netx

import (
	"net"
	"net/http"
)

const xForwardedFor = "X-Forwarded-For"

func GetRemoteIp(r *http.Request) string {
	v := r.Header.Get(xForwardedFor)
	if len(v) > 0 {
		return v
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}
