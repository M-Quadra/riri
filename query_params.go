package riri

import (
	"net/url"
)

// Params query_params
//   cover request line
type Params interface {
	Set(p map[string]string) *Request
	fullURL() string
}

// Set query params
func (r *Request) Set(p map[string]string) *Request {
	nValue := url.Values{}
	for k, v := range p {
		nValue.Set(k, v)
	}
	r.params = nValue
	return r
}

func (r Request) fullURL() string {
	if r.url == nil {
		return ""
	}

	fullURL := r.url.Scheme + "://" + r.url.Host + r.url.Path
	rawQuery := r.url.Query()
	if len(rawQuery) <= 0 {
		if len(r.params) <= 0 {
			return fullURL
		}

		return fullURL + "?" + r.params.Encode()
	}

	for k, v := range r.params {
		rawQuery[k] = v
	}

	return fullURL + "?" + rawQuery.Encode()
}
