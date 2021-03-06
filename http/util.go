package http

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
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

type Fetcher interface {
	FetchJson(interface{}) error
}

type fetcher struct {
	err error
	rc  io.ReadCloser
}

func (f fetcher) FetchJson(dst interface{}) error {
	if f.err != nil {
		return f.err
	}
	defer f.rc.Close()

	return json.NewDecoder(f.rc).Decode(dst)
}

func (f fetcher) FetchFile(file string) error {
	if f.err != nil {
		return f.err
	}
	defer f.rc.Close()

	data, err := ioutil.ReadAll(f.rc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0666)
}

func NewFetcher(spec *HttpSpec) (ftc fetcher) {
	var req *http.Request
	var rsp *http.Response
	var err error

	defer func() {
		ftc.err = err
		if rsp != nil && rsp.Body != nil {
			ftc.rc = rsp.Body
		}
	}()

	if req, err = NewRequest(spec); err != nil {
		return
	}

	rsp, err = (&http.Client{Transport: clientTransport}).Do(req)
	return
}
