package orch

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-resty/resty/v2"
)

func newHTTPError(statusCode int, msg string) *HTTPError {
	return &HTTPError{
		statusCode: statusCode,
		msg:        msg,
	}
}

// FormatError takes the resty response and a possible resty err and tries to create an
// OrchestratorError with as much info as possible
func FormatError(r *resty.Response, customError ...string) error {

	msg := strings.Join(customError, ", ")

	if orchErr, ok := r.Error().(*OrchestratorError); ok {
		if reflect.DeepEqual(&OrchestratorError{}, orchErr) {
			return newHTTPError(r.StatusCode(), msg)
		}

		orchErr.StatusCode = r.StatusCode()
		if len(msg) > 0 {
			orchErr.Msg = msg
		}

		return orchErr

	} else if r.IsError() {
		return newHTTPError(r.StatusCode(), msg)
	}

	return fmt.Errorf("%s", msg)
}
