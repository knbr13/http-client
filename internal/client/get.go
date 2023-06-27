package client

import (
	"net/http"
)

// Get sends an HTTP GET request.
func (c *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	// Create an HTTP request
	httpRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		httpRequest.Header.Set(key, value)
	}

	// Send the HTTP request
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	return httpResponse, nil
}
