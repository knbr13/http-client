package httpmethods

import (
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

func get(input Input) (*http.Response, error) {
	if input.HTTPMethod == "" || input.URL == "" {
		return nil, fmt.Errorf("missing some required arguments")
	}

	// Create an HTTP request
	httpRequest, err := http.NewRequest(http.MethodGet, input.URL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	headers, err := utilities.ParseHeaders(input.Header)
	if err != nil {
		return nil, err
	}

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
