package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gookit/color"
)

func printColoredHeaders(headers map[string][]string) {
	// Print headers with colors
	for key, values := range headers {
		for _, value := range values {
			color.Bold.Printf("%s: ", key)
			color.Cyan.Printf("%s\n", value)
		}
	}
}

// printColoredBody prints the colored JSON response body.
func printColoredBody(body []byte) {
	// Try parsing the body as a single JSON object
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err == nil {
		printColoredJSONObject(data)
		return
	}

	// Try parsing the body as an array of JSON objects
	var arrayData []map[string]interface{}
	err = json.Unmarshal(body, &arrayData)
	if err == nil {
		for _, obj := range arrayData {
			printColoredJSONObject(obj)
			fmt.Println()
		}
		return
	}

	// If parsing fails, print the raw body as plain text
	fmt.Println(string(body))
}

func printColoredJSONObject(obj map[string]interface{}) {
	keyColor := color.New(color.FgYellow)
	valueColor := color.New(color.FgCyan)

	for key, value := range obj {
		keyColor.Printf("%s: ", key)
		valueColor.Println(value)
	}
}

func writeToFile(fileName string, data []byte) error {
	if fileName == "" {
		return errors.New("missing file name")
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	datawriter := bufio.NewWriter(file)

	_, err = datawriter.WriteString(string(data))
	if err != nil {
		return err
	}

	color.Greenp("Saved to ", fileName, "\n")
	datawriter.Flush()

	return nil
}
