package client

import "net/http"

// Client represents the HTTP client.
type Client struct {
	httpClient *http.Client
	headers    http.Header
}
