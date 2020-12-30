package riri

// GET request
func GET(url string) Request {
	return newRequest("GET", url)
}
