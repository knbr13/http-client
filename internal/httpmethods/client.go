package httpmethods

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Command struct {
	Name  string
	Short string
	Run   func(Input) (*http.Response, error)
}

var AvailableHttpMethods = map[string]Command{
	"GET": {
		Name:  "GET",
		Short: "Sends a GET request to a specified URL",
		Run:   get,
	},
	"POST": {
		Name:  "POST",
		Short: "Sends a POST request to a specified URL",
		Run:   post,
	},
	"PUT": {
		Name:  "PUT",
		Short: "Sends a PUT request to a specified URL",
		Run:   put,
	},
	"PATCH": {
		Name:  "PATCH",
		Short: "Sends a PATCH request to a specified URL",
		Run:   patch,
	},
	"DELETE": {
		Name:  "DELETE",
		Short: "Sends a DELETE request to a specified URL",
		Run:   delete,
	},
	"HEAD": {
		Name:  "HEAD",
		Short: "Sends a HEAD request to a specified URL",
		Run:   head,
	},
	"OPTIONS": {
		Name:  "OPTIONS",
		Short: "Sends a OPTIONS request to a specified URL",
		Run:   options,
	},
}

type Input struct {
	HTTPMethod string `arg:"-m,--http-method,required"`
	URL        string `arg:"-u,--url"`
	Body       string `arg:"-b,--body"`
	Header     string `arg:"-H,--header"`
	Output     string `arg:"-o"`
}

var httpClient *http.Client = &http.Client{
	Timeout: 30 * time.Second,
}

func RunHttpMethod(input Input) (*http.Response, error) {
	command, ok := AvailableHttpMethods[strings.ToUpper(input.HTTPMethod)]
	if !ok {
		return nil, fmt.Errorf("unknown http method: %v", input.HTTPMethod)
	}

	return command.Run(input)
}
