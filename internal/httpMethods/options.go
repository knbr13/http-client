package httpmethods

import (
	"net/http"
)

// Options sends an HTTP OPTIONS request.
func Options(url string, body []byte, headers map[string]string) (*http.Response, error) {
	// Create a new HTTP client
	httpClient := &http.Client{}

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
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
