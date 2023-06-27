package internal

import "net/http"

type httpCommand struct {
	name        string
	description string
	run         func() (*http.Response, error)
}
