package httperr

import "fmt"

type HttpError interface {
	Error() string
	Code() int
	Message() string
}

type httpError struct {
	code    int
	message string
}

func (e *httpError) Code() int {
	return e.code
}

func (e *httpError) Message() string {
	return e.message
}

func (e *httpError) Error() string {
	return fmt.Sprintf("code=%d; message=%s", e.code, e.message)
}

func NewHttpError(code int, message string) HttpError {
	return &httpError{
		code:    code,
		message: message,
	}
}

type ReadableHttpError interface {
	HttpError
	ReadableMessage() string
}

type readableHttpError struct {
	httpError
	readableMessage string
}

func (e *readableHttpError) ReadableMessage() string {
	return e.readableMessage
}

func NewReadableHttpError(code int, message string, readableMessage string) ReadableHttpError {
	return &readableHttpError{
		httpError: httpError{
			code:    code,
			message: message,
		},
		readableMessage: readableMessage,
	}
}
