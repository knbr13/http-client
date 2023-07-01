package httpmethods

import (
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

// RunOptions executes the OPTIONS request using the provided input parameters.
func RunOptions(input []string) (*http.Response, error) {
	url, headers, err := parseOptionsInput(input)
	if err != nil {
		return nil, err
	}

	return options(url, headers)
}

// Options sends an HTTP OPTIONS request.
func options(url string, headers map[string]string) (*http.Response, error) {
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

// parseOptionsInput parses the input slice for the OPTIONS command and extracts the URL and headers.
func parseOptionsInput(input []string) (string, map[string]string, error) {
	if len(input) < 1 || len(input) > 2 {
		return "", nil, fmt.Errorf("invalid input for OPTIONS command")
	}

	url := input[0]
	headersStr := ""
	if len(input) == 2 {
		headersStr = input[1]
	}

	headers, err := utilities.ParseHeaders(headersStr)
	if err != nil {
		return "", nil, err
	}

	return url, headers, nil
}
