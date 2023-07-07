package httpmethods

import (
	"net/http"
	"time"
)

var AvailableHttpMethods = map[string]bool{
	"GET": true,
	"POST": true,
    "PUT": true,
	"PATCH": true,
    "DELETE": true,
	"HEAD": true,
    "OPTIONS": true,
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
