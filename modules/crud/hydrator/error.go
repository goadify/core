package hydrator

import (
	"github.com/goadify/goadify/modules/crud/httperr"
	"github.com/goadify/openapi/crud/go/gen"
	"net/http"
)

type errorHydrator struct {
	displayInternalMessages bool
}

type ErrorOption func(*errorHydrator)

func ErrorDisplayInternalMessages(display bool) ErrorOption {
	return func(hydrator *errorHydrator) {
		hydrator.displayInternalMessages = display
	}
}

func Error(err error, options ...ErrorOption) *gen.ErrorStatusCode {
	hydrator := new(errorHydrator)
	for _, option := range options {
		option(hydrator)
	}

	code := http.StatusInternalServerError
	message := http.StatusText(code)
	displayMessage := gen.OptString{}

	if hydrator.displayInternalMessages {
		message = err.Error()
	}

	if httpErr, ok := err.(httperr.ReadableHttpError); ok {
		code = httpErr.Code()
		message = httpErr.Message()
		displayMessage.SetTo(httpErr.ReadableMessage())
	} else if httpErr, ok := err.(httperr.HttpError); ok {
		code = httpErr.Code()
		message = httpErr.Message()
	}

	return &gen.ErrorStatusCode{
		StatusCode: code,
		Response: gen.Error{
			Code:           code,
			Message:        message,
			DisplayMessage: displayMessage,
		},
	}
}
