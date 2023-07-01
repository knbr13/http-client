package httpmethods

import (
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

// RunDelete executes the DELETE request using the provided input parameters.
func RunDelete(input []string) (*http.Response, error) {
	url, headers, err := parseDeleteInput(input)
	if err != nil {
		return nil, err
	}

	return delete(url, nil, headers)
}

// Delete sends an HTTP DELETE request.
func delete(url string, body []byte, headers map[string]string) (*http.Response, error) {
	// Create an HTTP request
	httpRequest, err := http.NewRequest(http.MethodDelete, url, nil)
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

// parseDeleteInput parses the input slice for the DELETE command and extracts the URL and headers.
func parseDeleteInput(input []string) (string, map[string]string, error) {
	if len(input) < 1 || len(input) > 2 {
		return "", nil, fmt.Errorf("invalid input for DELETE command")
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
