package orch

import (
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
)

func newHTTPError(statusCode int, msg string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Msg:        msg,
	}
}

// ProcessError will process the response for RESTY and thow an error accordingly
func ProcessError(r *resty.Response, err error, errorString string) error {
	// Return HTTP error detail if we have enough.
	if r.IsError() {
		var message string
		if len(errorString) > 0 {
			message = errorString
		} else {
			message = r.Status()
		}
		if orchErr, ok := r.Error().(*OrchestratorError); ok {
			if reflect.DeepEqual(&OrchestratorError{}, orchErr) {
				return newHTTPError(r.StatusCode(), message)
			}

			orchErr.StatusCode = r.StatusCode()
			return orchErr
		}
		return newHTTPError(r.StatusCode(), message)
	}

	// Cater for an error which didn't come from a HTTP response. (e.g. host not listening)
	if err != nil {
		if len(errorString) > 0 {
			return fmt.Errorf("%s", errorString)
		}
		return err
	}

	return nil
}
