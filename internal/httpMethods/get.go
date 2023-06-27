package httpmethods

import (
	"net/http"
)

// Get sends an HTTP GET request.
func Get(url string, body []byte, headers map[string]string) (*http.Response, error) {
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
	httpClient := &http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
