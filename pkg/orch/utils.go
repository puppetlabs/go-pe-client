package orch

import (
	"strings"

	"github.com/go-resty/resty/v2"
)

// UnmarshalError represents an error response from the Orchestrator Error Cast
type UnmarshalError struct {
	Msg        string
	StatusCode int
}

func (oe *UnmarshalError) Error() string {
	return oe.Msg
}

// FormatOrchError takes the resty response and a possible resty err and tries to create an
// OrchestratorError with as much info as possible
func FormatOrchError(r *resty.Response, customError ...string) error {

	orchErr, ok := r.Error().(*OrchestratorError)

	if !ok {
		return &UnmarshalError{
			Msg:        "unable to unmarshal an OrchestratorError",
			StatusCode: r.StatusCode(),
		}
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
