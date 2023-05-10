package types

import "net/http"

type Module interface {
	Name() string

	Init(CommonProvider) error

	HttpPrefix() string
	HttpHandler() (http.Handler, error)
}
