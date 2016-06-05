package http_client

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var pooledConfig = &tls.Config{
	InsecureSkipVerify: true,
	ClientSessionCache: tls.NewLRUClientSessionCache(128), // switch on cache of TLS sessions
}

var pooledTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       pooledConfig,
	MaxIdleConnsPerHost:   128, // expand connection pool size
}

var pooledClient = &http.Client{Transport: pooledTransport, Timeout: time.Second}
