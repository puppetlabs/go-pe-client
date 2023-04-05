package rbac

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRBACToken(t *testing.T) {
	// Test success
	setupPostResponder(t, requestAuthTokenURI, "GetRBACToken-request.json", "GetRBACToken-response.json")
	request := &RequestKeys{
		Login:    "jimbo",
		Password: "package",
	}

	expectedResponse := &Token{
		Token: "some token",
	}
	actual, err := rbacClient.GetRBACToken(request)
	require.Nil(t, err)
	require.Equal(t, expectedResponse, actual)

	// Test error
	setUpBadRequestResponder(t, http.MethodPost, requestAuthTokenURI)
	actual, err = rbacClient.GetRBACToken(request)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}
