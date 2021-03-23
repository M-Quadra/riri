package riri

// PreRequest url
type PreRequest string

// POST riri.POST
func (slf PreRequest) POST() Request {
	return POST(string(slf))
}

// GET riri.GET
func (slf PreRequest) GET() Request {
	return GET(string(slf))
}

// Group is used to generate adaptive url path in different environments.
type Group struct {
	BaseURL func() string
}

// Path return baseURL + path
func (slf Group) Path(path string) PreRequest {
	if slf.BaseURL == nil {
		return PreRequest(path)
	}

	return PreRequest(slf.BaseURL() + path)
}
