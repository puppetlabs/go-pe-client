package rbac

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-resty/resty/v2"
)

func newAPIError(statusCode int, msg string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Msg:        msg,
	}
}

// FormatError takes the resty response and a possible resty err and tries to create an
// APIError with as much info as possible
func FormatError(r *resty.Response, customError ...string) error {

	msg := strings.Join(customError, ", ")

	if apiErr, ok := r.Error().(*APIError); ok {
		if reflect.DeepEqual(&APIError{}, apiErr) {
			return newAPIError(r.StatusCode(), msg)
		}
		apiErr.StatusCode = r.StatusCode()
		if len(msg) > 0 {
			apiErr.Msg = msg
		}

		return apiErr

	} else if r.IsError() {
		return newAPIError(r.StatusCode(), msg)
	}

	return fmt.Errorf("%s", msg)
}
