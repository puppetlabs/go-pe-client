package orch

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// FormatOrchError takes the resty response and a possible resty err and tries to create an
// OrchestratorError with as much info as possible
func FormatOrchError(r *resty.Response, customError ...string) error {

	orchErr, ok := r.Error().(*OrchestratorError)

	if !ok {
		return fmt.Errorf("unable to unmarshal an OrchestratorError, StatusCode: %s", r.Status())
	}

	if len(customError) > 0 {
		orchErr.Msg = strings.Join(customError, ", ")
	}

	x := &OrchestratorError{
		StatusCode: r.StatusCode(),
		Kind:       orchErr.Kind,
		Msg:        orchErr.Msg,
	}

	return x
}
