package httpmethods

import (
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

// RunHead executes the HEAD request using the provided input parameters.
func RunHead(input []string) (*http.Response, error) {
	url, headers, err := parseHeadInput(input)
	if err != nil {
		return nil, err
	}

	return head(url, headers)
}

// Head sends an HTTP HEAD request.
func head(url string, headers map[string]string) (*http.Response, error) {
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
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

// parseHeadInput parses the input slice for the HEAD command and extracts the URL and headers.
func parseHeadInput(input []string) (string, map[string]string, error) {
	if len(input) < 1 || len(input) > 2 {
		return "", nil, fmt.Errorf("invalid input for HEAD command")
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
