package rbac

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	getRoleResponseFilePath  = "testdata/apidocs/GetRole-response.json"
	getRolesResponseFilePath = "testdata/apidocs/GetRoles-response.json"
	token                    = "dummy-token"
)

func TestGetRoles(t *testing.T) {
	var expectedRoles []Role

	expectedRolesJSONFile, err := os.Open(getRolesResponseFilePath)
	require.Nil(t, err, "failed to open expected roles JSON file")

	err = json.NewDecoder(expectedRolesJSONFile).Decode(&expectedRoles)
	require.Nil(t, err, "error decoding expected roles")

	setUpOKResponder(t, http.MethodGet, rolesPath, getRolesResponseFilePath)

	actualRoles, err := rbacClient.GetRoles(token)
	require.Nil(t, err)
	require.Equal(t, expectedRoles, actualRoles)
}

func TestGetRole(t *testing.T) {
	var expectedRole *Role

	expectedRoleJSONFile, err := os.Open(getRoleResponseFilePath)
	require.Nil(t, err, "failed to open expected role JSON file")

	err = json.NewDecoder(expectedRoleJSONFile).Decode(&expectedRole)
	require.Nil(t, err, "error decoding expected role")

	rolePathWithID := strings.ReplaceAll(rolePath, "{id}", strconv.Itoa(int(expectedRole.ID)))
	setUpOKResponder(t, http.MethodGet, rolePathWithID, getRoleResponseFilePath)

	actualRole, err := rbacClient.GetRole(expectedRole.ID, token)
	require.Nil(t, err)
	require.Equal(t, expectedRole, actualRole)
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
