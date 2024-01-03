package utilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseHeaders parses the input headers string into a map.
func ParseHeaders(headersStr string) (map[string]string, error) {
	headers := make(map[string]string)
	headersStr = strings.TrimSpace(headersStr)

	if len(headersStr) == 0 {
		return headers, nil
	}

	var headerMap map[string]string
	err := json.Unmarshal([]byte(headersStr), &headerMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse headers: %v", err)
	}

	return headerMap, nil
}

// ParseBody parses the input body string into a JSON byte slice.
func ParseBody(bodyStr string) ([]byte, error) {
	bodyStr = strings.TrimSpace(bodyStr)

	if len(bodyStr) == 0 {
		return []byte{}, nil
	}

	if !strings.HasPrefix(bodyStr, "{") || !strings.HasSuffix(bodyStr, "}") {
		return nil, fmt.Errorf("invalid body format: %s", bodyStr)
	}

	bodyStr = strings.TrimPrefix(bodyStr, "{")
	bodyStr = strings.TrimSuffix(bodyStr, "}")

	keyValuePairs := strings.Split(bodyStr, ",")

	data := make(map[string]interface{})

	for _, kvPair := range keyValuePairs {
		parts := strings.Split(kvPair, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid body format")
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		data[key] = value
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return jsonData, nil
}
