package client

// Request represents an HTTP request.
type Request struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string]string
}
