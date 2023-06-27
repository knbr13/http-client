package client

import (
	"net/http"
)

// Head sends an HTTP HEAD request.
func (c *Client) Head(url string, headers map[string]string) (http.Header, error) {
	// Create an HTTP request
	httpRequest, err := http.NewRequest(http.MethodHead, url, nil)
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

	return httpResponse.Header, nil
}
