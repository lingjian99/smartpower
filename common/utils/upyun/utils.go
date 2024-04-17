package upyun

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"
)

func makeRFC1123Date(d time.Time) string {
	utc := d.UTC().Format(time.RFC1123)
	return strings.Replace(utc, "UTC", "GMT", -1)
}

func makeUserAgent(version string) string {
	return fmt.Sprintf("UPYUN Go SDK V2/%s", version)
}

func md5Str(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func base64ToStr(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func hmacSha1(key string, data []byte) []byte {
	hm := hmac.New(sha1.New, []byte(key))
	hm.Write(data)
	return hm.Sum(nil)
}

func escapeUri(uri string) (string, error) {
	uri = path.Join("/", uri)
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
