package client

import (
	"bytes"
	"io"
	"net/http"
)

// Post sends an HTTP POST request.
func (c *Client) Post(url string, body []byte, headers map[string]string) ([]byte, error) {
	// Create an HTTP request with the request body
	httpRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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
