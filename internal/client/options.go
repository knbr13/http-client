package client

import (
	"io"
	"net/http"
)

// Options sends an HTTP OPTIONS request.
func (c *Client) Options(url string, headers map[string]string) ([]byte, error) {
	// Create an HTTP request
	httpRequest, err := http.NewRequest(http.MethodOptions, url, nil)
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

	// Read the response body
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
