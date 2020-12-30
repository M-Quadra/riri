package riri

// POST request
func POST(url string) Request {
	return newRequest("POST", url)
}
