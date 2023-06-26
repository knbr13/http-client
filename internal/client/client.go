package client

import "net/http"

// Client represents the HTTP client.
type Client struct {
	httpClient *http.Client
	headers    http.Header
}

func NewHTTPClient() *Client {
	httpClient := &http.Client{}
	headers := make(http.Header)

	return &Client{
		httpClient: httpClient,
		headers:    headers,
	}
}
