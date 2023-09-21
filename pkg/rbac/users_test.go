package rbac

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	getUserResponseFilePath  = "testdata/apidocs/GetUser-response.json"
	getUsersResponseFilePath = "testdata/apidocs/GetUsers-response.json"
)

func TestGetUsers(t *testing.T) {
	var expectedUsers []User

	expectedUsersJSONFile, err := os.Open(getUsersResponseFilePath)
	require.Nil(t, err, "failed to open expected users JSON file")

	err = json.NewDecoder(expectedUsersJSONFile).Decode(&expectedUsers)
	require.Nil(t, err, "error decoding expected users")

	setUpOKResponder(t, requestUsersURI, getUsersResponseFilePath)

	actualUsers, err := rbacClient.GetUsers(token)
	require.Nil(t, err)
	require.NotNil(t, actualUsers)

	require.Equal(t, len(expectedUsers), len(actualUsers))
}

func TestGetCurrentUser(t *testing.T) {
	var expectedUser User

	expectedUserJSONFile, err := os.Open(getUserResponseFilePath)
	require.Nil(t, err, "failed to open expected user JSON file")

	err = json.NewDecoder(expectedUserJSONFile).Decode(&expectedUser)
	require.Nil(t, err, "error decoding expected role")

	setUpOKResponder(t, requestCurrentUserURI, getUserResponseFilePath)

	actualUser, err := rbacClient.GetCurrentUser(token)
	require.Nil(t, err)

	match := reflect.DeepEqual(expectedUser, *actualUser)
	require.True(t, match, "Expected and actual output do not match.")
}

func TestGetSpecificUser(t *testing.T) {
	specificUserID := "specific-user-test"

	var expectedUser User

	expectedUserJSONFile, err := os.Open(getUserResponseFilePath)
	require.Nil(t, err, "failed to open expected user JSON file")

	err = json.NewDecoder(expectedUserJSONFile).Decode(&expectedUser)
	require.Nil(t, err, "error decoding expected role")

	specificURI := fmt.Sprintf("%s%s", requestUserURI, specificUserID)

	setUpOKResponder(t, specificURI, getUserResponseFilePath)

	actualUser, err := rbacClient.GetSpecificUser(token, specificUserID)
	require.Nil(t, err)

	match := reflect.DeepEqual(expectedUser, *actualUser)
	require.True(t, match, "Expected and actual output do not match.")
}
