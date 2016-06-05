package http_client

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var defaultConfig = &tls.Config{
	InsecureSkipVerify: true,
}

var defaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       defaultConfig,
}

var defaultClient = &http.Client{Transport: defaultTransport, Timeout: time.Second}
