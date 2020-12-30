package riri

import "strings"

// Params query_params
//   cover request line
type Params interface {
	Set(p map[string]string) *Request
	fullURL() string
}

// Set query params
func (slf *Request) Set(p map[string]string) *Request {
	slf.params = p
	return slf
}

func (slf Request) fullURL() string {
	if len(slf.params) <= 0 {
		return slf.url
	}

	pAry := []string{}
	for k, v := range slf.params {
		if len(k) <= 0 {
			continue
		}

		if len(v) <= 0 {
			pAry = append(pAry, k)
			continue
		}

		pAry = append(pAry, k+"="+v)
	}
	if len(pAry) <= 0 {
		return slf.url
	}

	// request line
	ary := strings.Split(slf.url, "?")
	if len(ary) >= 2 {
		lineAry := strings.Split(ary[1], "&")
		for _, v := range lineAry {
			kv := strings.SplitN(v, "=", 2)
			_, ok := slf.params[kv[0]]
			if ok {
				continue
			}

			pAry = append(pAry, v)
		}
	}

	return ary[0] + "?" + strings.Join(pAry, "&")
}
