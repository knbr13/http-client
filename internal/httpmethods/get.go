package httpmethods

import (
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

// RunGet executes the GET request using the provided input parameters.
func RunGet(input []string) (*http.Response, error) {
	url, headers, err := parseGetInput(input)
	if err != nil {
		return nil, err
	}

	return get(url, nil, headers)
}

// Get sends an HTTP GET request.
func get(url string, body []byte, headers map[string]string) (*http.Response, error) {
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

// parseGetInput parses the input slice for the GET command and extracts the URL and headers.
func parseGetInput(input []string) (string, map[string]string, error) {
	if len(input) < 1 || len(input) > 2 {
		return "", nil, fmt.Errorf("invalid input for GET command")
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
