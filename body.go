package riri

import (
	"bytes"
	"encoding/json"

	"github.com/M-Quadra/kazaana"
)

// Body body
type Body interface {
	Binary(data []byte, contentType ...string) *Request
	JSON(v interface{}) *Request
}

// Binary body binary
func (slf *Request) Binary(data []byte, contentType ...string) *Request {
	slf.payload = bytes.NewReader(data)

	if contentType == nil {
		return slf
	}

	if slf.headers == nil {
		slf.headers = map[string]string{}
	}
	slf.headers["Content-Type"] = contentType[0]
	return slf
}

// JSON body raw JSON
func (slf *Request) JSON(v interface{}) *Request {
	slf.payload = nil
	slf.kerr = kazaana.New(nil)
	if v == nil {
		return slf
	}

	jsonData, err := json.Marshal(v)
	if err != nil {
		slf.kerr = kazaana.New(err)
		return slf
	}

	if slf.headers == nil {
		slf.headers = map[string]string{}
	}
	slf.headers["Content-Type"] = "application/json"
	slf.payload = bytes.NewReader(jsonData)
	return slf
}
