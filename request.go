package riri

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/M-Quadra/kazaana/v2"
)

// Request base
type Request struct {
	Params
	Headers
	Body

	method string
	url    *url.URL

	params  url.Values
	headers map[string]string
	payload *bytes.Reader

	kerr kazaana.Error
}

func newRequest(method, rawurl string) Request {
	u, err := url.Parse(rawurl)
	req := Request{
		method: method,
		url:    u,
	}
	if err != nil {
		req.kerr = kazaana.New(err)
	}

	req.Params = &req
	req.Headers = &req
	req.Body = &req
	return req
}

// Do client.Do
func (slf Request) Do() (*http.Response, kazaana.Error) {
	if slf.kerr.CheckError() {
		return nil, slf.kerr
	}

	var (
		req *http.Request
		err error
	)

	if slf.payload == nil {
		req, err = http.NewRequest(slf.method, slf.fullURL(), nil)
	} else {
		req, err = http.NewRequest(slf.method, slf.fullURL(), slf.payload)
	}
	if err != nil {
		return nil, kazaana.New(err)
	}

	for k, v := range slf.headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	return res, kazaana.New(err)
}

// Result res body data
func (slf Request) Result() ([]byte, kazaana.Error) {
	res, kerr := slf.Do()
	if kerr.CheckError() {
		return nil, kerr
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, kazaana.New(err)
}

// BindJSON .Result() -> JSON
//   without json.Number
func (slf Request) BindJSON(ptr interface{}) ([]byte, kazaana.Error) {
	body, kerr := slf.Result()
	if kerr.CheckError() {
		return nil, kerr
	}

	err := json.Unmarshal(body, ptr)
	return body, kazaana.New(err)
}
