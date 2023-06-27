package client

import (
	"bytes"
	"net/http"
)

// Put sends an HTTP PUT request.
func (c *Client) Put(url string, body []byte, headers map[string]string) (*http.Response, error) {
	// Create an HTTP request with the request body
	httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
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
