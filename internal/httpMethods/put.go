package httpmethods

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"

)

// Put sends an HTTP PUT request.
func put(url string, body []byte, headers map[string]string) (*http.Response, error) {
	// Create a new HTTP client
	httpClient := &http.Client{}

	// Create an HTTP request with the request body
	httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
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

// parsePutInput parses the input slice for the PUT command and extracts the URL, body, and headers.
func parsePutInput(input []string) (string, []byte, map[string]string, error) {
	if len(input) < 2 || len(input) > 3 {
		return "", nil, nil, fmt.Errorf("invalid input for PUT command")
	}

	url := input[0]
	bodyStr := input[1]
	headersStr := ""
	if len(input) == 3 {
		headersStr = input[2]
	}

	body, err := utilities.ParseBody(bodyStr)
	if err != nil {
		return "", nil, nil, err
	}

	headers, err := utilities.ParseHeaders(headersStr)
	if err != nil {
		return "", nil, nil, err
	}

	return url, body, headers, nil
}
