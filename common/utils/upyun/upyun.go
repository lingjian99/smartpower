package upyun

import (
	"fmt"
	"path"
	"strings"
	"time"
)

var version = "2.1.0"

type UpYunConfig struct {
	Bucket    string
	Operator  string
	Password  string
	Host      string `json:",optional"`
	UserAgent string `json:",optional"`
}

type UnifiedAuthConfig struct {
	Method     string
	Uri        string
	DateStr    string
	Policy     string
	ContentMD5 string
}

type UpYun struct {
	UpYunConfig
	uri string
	md5 string
}

func NewUpYun(config *UpYunConfig, path, md5 string) *UpYun {
	up := &UpYun{}
	up.Bucket = config.Bucket
	up.Operator = config.Operator
	up.Password = md5Str(config.Password)
	up.Host = config.Host
	if config.UserAgent != "" {
		up.UserAgent = config.UserAgent
	} else {
		up.UserAgent = makeUserAgent(version)
	}
	up.uri = path
	up.md5 = md5
	return up
}

func (up *UpYun) MakeHeaders() (headers map[string]string, url string, err error) {
	headers = make(map[string]string)

	escUri, err := escapeUri(up.uri)
	if err != nil {
		return
	}
	escUri = path.Join("/", up.Bucket, escUri)
	if strings.HasSuffix(up.uri, "/") {
		escUri += "/"
	}

	if len(up.md5) > 0 {
		headers["Content-MD5"] = up.md5
	}

	headers["Date"] = makeRFC1123Date(time.Now())
	headers["Host"] = "v0.api.upyun.com"

	headers["Authorization"] = up.MakeUnifiedAuth(&UnifiedAuthConfig{
		Method:     "PUT",
		Uri:        escUri,
		DateStr:    headers["Date"],
		ContentMD5: headers["Content-MD5"],
	})

	url = fmt.Sprintf("http://%s%s", up.Host, escUri)
	return
}

func (u *UpYun) MakeUnifiedAuth(config *UnifiedAuthConfig) string {
	sign := []string{
		config.Method,
		config.Uri,
		config.DateStr,
		config.Policy,
		config.ContentMD5,
	}
	var signNoEmpty []string
	for _, v := range sign {
		if v != "" {
			signNoEmpty = append(signNoEmpty, v)
		}
	}
	signStr := base64ToStr(hmacSha1(u.Password, []byte(strings.Join(signNoEmpty, "&"))))
	return "UPYUN " + u.Operator + ":" + signStr
}
