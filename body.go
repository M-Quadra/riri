package riri

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/M-Quadra/kazaana/v2"
)

// Body body
type Body interface {
	FormData(f map[string]FormDataValue) *Request
	JSON(v interface{}) *Request
	Binary(data []byte, contentType ...string) *Request
}

// Binary body binary
func (r *Request) Binary(data []byte, contentType ...string) *Request {
	r.payload = bytes.NewReader(data)

	if contentType == nil {
		return r
	}

	if r.headers == nil {
		r.headers = map[string]string{}
	}
	if len(contentType) > 0 {
		r.headers["Content-Type"] = contentType[0]
	}
	return r
}

// JSON body raw JSON
func (r *Request) JSON(v interface{}) *Request {
	jsonData, err := json.Marshal(v)
	if err != nil {
		r.kerr = kazaana.New(err)
		return r
	}

	return r.Body.Binary(jsonData, "application/json")
}

// FormDataValue type Text if len(FileData)<=0 || len(FileName)<=0
type FormDataValue struct {
	Text string

	FileName string
	FileData []byte
}

// FormData body raw form-data
func (r *Request) FormData(f map[string]FormDataValue) *Request {
	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)

	for k, v := range f {
		if len(v.FileData) <= 0 || len(v.FileName) <= 0 { // Text
			err := writer.WriteField(k, v.Text)
			if err != nil {
				r.kerr = kazaana.New(err)
				return r
			}
			continue
		}

		// File
		f, err := writer.CreateFormFile(k, v.FileName)
		if err != nil {
			r.kerr = kazaana.New(err)
			return r
		}

		_, err = io.Copy(f, bytes.NewReader(v.FileData))
		if err != nil {
			r.kerr = kazaana.New(err)
			return r
		}
	}

	err := writer.Close()
	if err != nil {
		r.kerr = kazaana.New(err)
		return r
	}

	data, err := io.ReadAll(formData)
	if err != nil {
		r.kerr = kazaana.New(err)
		return r
	}

	return r.Binary(data, writer.FormDataContentType())
}
