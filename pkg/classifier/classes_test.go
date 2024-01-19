package classifier

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

var (
	client        *Client
	classifierURL = "https://test-host:4433"
)

func init() {
	client = NewClient(classifierURL, "xxxx", nil)
	httpmock.Activate()
	httpmock.ActivateNonDefault(client.resty.GetClient())
}

func TestClassesSuccess(t *testing.T) {
	setupGetResponder(t, classes, "", "classes-response.json")
	actual, err := client.Classes(nil)

	expected := []Class{{Name: "classinproduction", Environment: "production"},
		{Name: "classindev", Environment: "development"}}

	require.Nil(t, err)
	require.True(t, len(actual) == len(expected))
	require.Equal(t, expected, actual)
}

func TestClassesFailure(t *testing.T) {
	setupResponderWithStatusCodeAndBody(t, classes, http.StatusInternalServerError, nil)
	_, err := client.Classes(nil)

	require.NotNil(t, err)
}
