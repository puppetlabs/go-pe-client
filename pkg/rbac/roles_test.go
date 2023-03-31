package rbac

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	responseFilePath = "testdata/apidocs/GetRoles-response.json"
	token            = "dummy-token"
)

func TestGetRoles(t *testing.T) {
	var expectedRoles []Role

	expectedRolesJSONFile, err := os.Open(responseFilePath)
	require.Nil(t, err, "failed to open expected roles JSON file")

	err = json.NewDecoder(expectedRolesJSONFile).Decode(&expectedRoles)
	require.Nil(t, err, "error decoding expected roles")

	setUpOKResponder(t, http.MethodGet, rolesPath, responseFilePath)

	actualRoles, err := rbacClient.GetRoles(token)
	require.Nil(t, err)
	require.Equal(t, expectedRoles, actualRoles)
}

func TestCreateRole(t *testing.T) {
	role := &Role{
		DisplayName: "Testing",
		Description: "Role added by go-pe-client test",
		Permissions: []Permission{
			{
				ObjectType: "node_groups",
				Action:     "view",
				Instance:   "*",
			},
		},
	}

	// Test success
	setupCreateRoleSuccessResponder(t, rolesPath, "CreateRole-request.json")
	actual, err := rbacClient.CreateRole(role, token)
	require.Nil(t, err)
	require.Equal(t, "/path/to/role", actual)

	// Test error
	setupCreateRoleErrorResponder(t, rolesPath)
	location, err := rbacClient.CreateRole(role, token)
	require.NotNil(t, err)
	require.Equal(t, "", location)
	require.Equal(t, 409, err.(*APIError).GetStatusCode())
	require.Contains(t, err.(*APIError).Error(), "database conflict")
}
