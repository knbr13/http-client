package httpmethods

import (
	"net/http"
	"time"
)

type Command struct {
	Name  string
	Short string
	Run   func(Input) (*http.Response, error)
}

var AvailableHttpMethods = map[string]Command{
	"GET": Command{
		Name:  "GET",
		Short: "Sends a GET request to a specified URL",
		Run:   get,
	},
	"POST": Command{
		Name:  "POST",
		Short: "Sends a POST request to a specified URL",
		Run:   post,
	},
	"PUT": Command{
		Name:  "PUT",
		Short: "Sends a PUT request to a specified URL",
		Run:   put,
	},
	"PATCH": Command{
		Name:  "PATCH",
		Short: "Sends a PATCH request to a specified URL",
		Run:   patch,
	},
	"DELETE": Command{
		Name:  "DELETE",
		Short: "Sends a DELETE request to a specified URL",
		Run:   delete,
	},
	"HEAD": Command{
		Name:  "HEAD",
		Short: "Sends a HEAD request to a specified URL",
		Run:   head,
	},
	"OPTIONS": Command{
		Name:  "OPTIONS",
		Short: "Sends a OPTIONS request to a specified URL",
		Run:   options,
	},
}

type Input struct {
	HTTPMethod string `arg:"-m,--http-method"`
	URL        string `arg:"-u,--url"`
	Body       string `arg:"-b,--body"`
	Header     string `arg:"-H,--header"`
}

var httpClient *http.Client = &http.Client{
	Timeout: 30 * time.Second,
}
