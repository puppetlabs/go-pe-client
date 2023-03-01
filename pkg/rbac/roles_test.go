package rbac

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRole(t *testing.T) {

	var (
		token = "dummyToken"

		role = &Role{
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
	)

	// Test success
	setupCreateRoleSuccessResponder(t, rbacRoles, "CreateRole-request.json")
	actual, err := rbacClient.CreateRole(role, token)
	require.Nil(t, err)
	require.Equal(t, "/path/to/role", actual)

	// Test error
	setupCreateRoleErrorResponder(t, rbacRoles)
	location, err := rbacClient.CreateRole(role, token)
	require.NotNil(t, err)
	require.Equal(t, "", location)
	require.Equal(t, 409, err.(*APIError).GetStatusCode())
	require.Contains(t, err.(*APIError).Error(), "database conflict")
}
