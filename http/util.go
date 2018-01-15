package http

import (
	"time"
	"net/http"
	"crypto/tls"
)

var clientTransport = &http.Transport{
	MaxIdleConns:        500,
	MaxIdleConnsPerHost: 100,
	IdleConnTimeout:     time.Minute * 10,
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

func Send(spec *HttpSpec) (*http.Response, error) {
	req, err := NewRequest(spec)
	if err != nil {
		return nil, err
	}
	cli := http.Client{
		Transport: clientTransport,
	}
	return cli.Do(req)
}
