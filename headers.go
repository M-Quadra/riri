package riri

// Headers headers
type Headers interface {
	Add(h map[string]string) *Request
}

// Add headers
func (slf *Request) Add(h map[string]string) *Request {
	slf.headers = h
	return slf
}
