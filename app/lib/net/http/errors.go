package http

import (
	"net/http"
)

func AbortHandler() {
	panic(http.ErrAbortHandler)
}

func IsHandlerAborted(v interface{}) bool {
	return v == http.ErrAbortHandler
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

type ErrHttpStatus struct {
	Err    error
	Status int
}

func (e *ErrHttpStatus) Error() string {
	return e.Err.Error()
}
