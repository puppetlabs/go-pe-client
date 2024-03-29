package rbac

import (
	"fmt"
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

func TestAuthenticateRBACToken(t *testing.T) {
	// Test success
	setupPostResponder(t, tokenAuthenticateURI, "AuthenticateRBACToken-request.json",
		"AuthenticateRBACToken-response.json")

	expectedResponse := &AuthenticateResponse{
		UserID:      "abc",
		Description: "abc",
		RoleIDs:     []int{1, 2, 3},
	}
	actual, err := rbacClient.AuthenticateRBACToken("blah")
	require.Nil(t, err)
	require.Equal(t, expectedResponse, actual)

	// Test error
	setUpBadRequestResponder(t, http.MethodPost, tokenAuthenticateURI)
	actual, err = rbacClient.AuthenticateRBACToken("blah")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)
}

func TestRevokeRBACToken(t *testing.T) {
	tokenValue := "abc"

	// Test success
	setUpOKDeleteResponder(fmt.Sprintf("%s%s", tokenRevokeURI, tokenValue))

	err := rbacClient.RevokeRBACToken(tokenValue)
	require.Nil(t, err)

	// Test error
	setUpBadRequestResponder(t, http.MethodDelete, fmt.Sprintf("%s%s", tokenRevokeURI, tokenValue))
	err = rbacClient.RevokeRBACToken(tokenValue)
	require.Equal(t, expectedError, err)
}

func TestGenerateRBACToken(t *testing.T) {
	tokenValue := "some token"

	tokenRequest := TokenRequest{
		Description: "A token to be used with joy and care.",
		Lifetime:    "1y",
		Client:      "PE console",
	}

	// Test success
	setupPostResponder(t, tokenGenerateURI, "GenerateToken-request.json", "GenerateToken-response.json")

	tokenResponse, err := rbacClient.GenerateRBACToken(tokenValue, tokenRequest)
	require.Nil(t, err)
	require.Equal(t, tokenValue, tokenResponse)

	// Test error
	setUpBadRequestResponder(t, http.MethodPost, tokenGenerateURI)
	_, err = rbacClient.GenerateRBACToken(tokenValue, tokenRequest)
	require.Equal(t, expectedError, err)
}
