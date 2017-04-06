package utils

import (
	"crypto/tls"
	"net/http"
)

//解决 x509 未认证的验证问题
func NewClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}