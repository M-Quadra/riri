package riri

// URLGroup is used to generate adaptive url path in different environments.
type URLGroup struct {
	BaseURL func() string
}

// Path return baseURL + path
func (slf *URLGroup) Path(path string) string {
	if slf.BaseURL == nil {
		return path
	}

	return slf.BaseURL() + path
}
