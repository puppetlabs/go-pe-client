package orch

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTasks(t *testing.T) {
	// Test without environment
	setupGetResponder(t, orchTasks, "", "tasks-response.json")
	actual, err := orchClient.Tasks("")
	require.Nil(t, err)
	require.Equal(t, expectedTasks, actual)

	// Test with environment
	setupGetResponder(t, orchTasks, "environment=myenv", "tasks-response.json")
	actual, err = orchClient.Tasks("myenv")
	require.Nil(t, err)
	require.Equal(t, expectedTasks, actual)

	// Test error
	setupErrorResponder(t, orchTasks)
	actual, err = orchClient.Tasks("")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

	// Test HTTP error
	setupResponderWithStatusCodeAndBody(t, orchTasks, http.StatusNotFound, []byte(`{"StatusCode": 400}`))
	actual, err = orchClient.Tasks("")
	testHTTPError(t, actual, err, http.StatusNotFound)
}

func TestTask(t *testing.T) {
	replacer := strings.NewReplacer("{module}", "foo", "{taskname}", "bar")
	orchTaskFooBar := replacer.Replace(orchTaskName)

	// Test without environment
	setupGetResponder(t, orchTaskFooBar, "", "task-response.json")
	actual, err := orchClient.Task("", "foo", "bar")
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test with environment
	setupGetResponder(t, orchTaskFooBar, "environment=myenv", "task-response.json")
	actual, err = orchClient.Task("myenv", "foo", "bar")
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test error
	setupErrorResponder(t, orchTaskFooBar)
	actual, err = orchClient.Task("myenv", "foo", "bar")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

	// Test HTTP error
	setupResponderWithStatusCodeAndBody(t, orchTaskFooBar, http.StatusNotFound, []byte(`{"StatusCode": 400}`))
	actual, err = orchClient.Task("myenv", "foo", "bar")
	testHTTPError(t, actual, err, http.StatusNotFound)
}

func TestTaskByID(t *testing.T) {
	replacer := strings.NewReplacer("{module}", "package", "{taskname}", "upgrade")
	orchTaskPackageUpgrade := replacer.Replace(orchTaskName)

	id := "https://orchestrator.example.com:8143" + orchTaskPackageUpgrade

	// Test without environment
	setupGetResponder(t, orchTaskPackageUpgrade, "", "task-response.json")
	actual, err := orchClient.TaskByID("", id)
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test with environment
	setupGetResponder(t, orchTaskPackageUpgrade, "environment=myenv", "task-response.json")
	actual, err = orchClient.TaskByID("myenv", id)
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test error
	setupErrorResponder(t, orchTaskPackageUpgrade)
	actual, err = orchClient.TaskByID("myenv", id)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

	// Test HTTP error
	setupResponderWithStatusCodeAndBody(t, orchTaskPackageUpgrade, http.StatusNotFound, []byte(`{"StatusCode": 400}`))
	actual, err = orchClient.TaskByID("myenv", id)
	testHTTPError(t, actual, err, http.StatusNotFound)
}

var expectedTask = &Task{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/install", Name: "package::install", Environment: struct {
	Name   string "json:\"name,omitempty\""
	CodeID string "json:\"code_id,omitempty\""
}{Name: "production", CodeID: "urn:puppet:code-id:1:a86da166c30f871823f9b2ea224796e834840676;production"}, Metadata: TaskMetadata{Description: "Bootstrap a node with puppet-agent", Private: true, SupportsNoop: false, InputMethod: "stdin", Parameters: map[string]TaskParam{"cacert_content": {Description: "The expected CA certificate content for the master", Type: "Optional[String]"}, "certname": {Description: "The certname with which the node should be bootstrapped", Type: "Optional[String]"}, "custom_attribute": {Description: "This setting is added to puppet.conf and included in the custom_attributes section of csr_attributes.yaml", Type: "Optional[Array[Pattern[/\\w+=\\w+/]]]"}, "dns_alt_names": {Description: "The DNS alt names with which the agent certificate should be generated", Type: "Optional[String]"}, "environment": {Description: "The environment in which the node should be bootstrapped", Type: "Optional[String]"}, "extension_request": {Description: "This setting is added to puppet.conf and included in the extension_requests section of csr_attributes.yaml", Type: "Optional[Array[Pattern[/\\w+=\\w+/]]]"}, "master": {Description: "The fqdn of the master from which the puppet-agent should be bootstrapped", Type: "String"}, "set_noop": {Description: "The noop setting in the [agent] section of puppet.conf", Type: "Optional[Boolean]"}}, Extensions: map[string]interface{}{"discovery": map[string]interface{}{"friendlyName": "Install Puppet agent", "parameters": map[string]interface{}{"cacert_content": map[string]interface{}{"placeholder": "-----BEGIN CERTIFICATE---- ... -----END CERTIFICATE-----"}, "master": map[string]interface{}{"placeholder": "master.company.net"}}, "puppetInstall": true, "type": []interface{}{"host"}}}, Implementations: []TaskImplementation{{Name: "windows.ps1", Requirements: []string{"powershell"}, InputMethod: "powershell"}, {Name: "linux.sh", Requirements: []string{"shell"}, InputMethod: "environment"}}}, Files: []TaskFile{{Filename: "install", URI: struct {
	Path   string "json:\"path,omitempty\""
	Params struct {
		Environment string "json:\"environment,omitempty\""
	} "json:\"params,omitempty\""
}{Path: "/package/tasks/install", Params: struct {
	Environment string "json:\"environment,omitempty\""
}{Environment: "production"}}, Sha256: "a9089b5b9720dca38a49db6f164cf8a053a7ea528711325da1c23de94672980f", SizeBytes: 693}}}

var expectedTasks = &Tasks{Environment: struct {
	Name   string "json:\"name,omitempty\""
	CodeID string "json:\"code_id,omitempty\""
}{Name: "production", CodeID: "urn:puppet:code-id:1:a86da166c30f871823f9b2ea224796e834840676;production"}, Items: []struct {
	ID   string "json:\"id,omitempty\""
	Name string "json:\"name,omitempty\""
}{
	{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/install", Name: "package::install"},

	{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/upgrade", Name: "package::upgrade"},

	{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/exec/init", Name: "exec"}}}
