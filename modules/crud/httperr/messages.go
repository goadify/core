package httperr

import "net/http"

var (
	InternalServerErrorMessage = http.StatusText(http.StatusInternalServerError)
)
