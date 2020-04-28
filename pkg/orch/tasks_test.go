package orch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTasks(t *testing.T) {

	// Test without environment
	setupGetResponder(t, "/orchestrator/v1/tasks", "", "tasks-response.json")
	actual, err := orchClient.Tasks("")
	require.Nil(t, err)
	require.Equal(t, expectedTasks, actual)

	// Test with environment
	setupGetResponder(t, "/orchestrator/v1/tasks", "environment=myenv", "tasks-response.json")
	actual, err = orchClient.Tasks("myenv")
	require.Nil(t, err)
	require.Equal(t, expectedTasks, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/tasks")
	actual, err = orchClient.Tasks("")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestTask(t *testing.T) {

	// Test without environment
	setupGetResponder(t, "/orchestrator/v1/tasks/foo/bar", "", "task-response.json")
	actual, err := orchClient.Task("", "foo", "bar")
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test with environment
	setupGetResponder(t, "/orchestrator/v1/tasks/foo/bar", "environment=myenv", "task-response.json")
	actual, err = orchClient.Task("myenv", "foo", "bar")
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/tasks/foo/bar")
	actual, err = orchClient.Task("myenv", "foo", "bar")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestTaskByID(t *testing.T) {

	id := "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/upgrade"

	// Test without environment
	setupGetResponder(t, "/orchestrator/v1/tasks/package/upgrade", "", "task-response.json")
	actual, err := orchClient.TaskByID("", id)
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test with environment
	setupGetResponder(t, "/orchestrator/v1/tasks/package/upgrade", "environment=myenv", "task-response.json")
	actual, err = orchClient.TaskByID("myenv", id)
	require.Nil(t, err)
	require.Equal(t, expectedTask, actual)

	// Test error
	setupErrorResponder(t, "/orchestrator/v1/tasks/package/upgrade")
	actual, err = orchClient.TaskByID("myenv", id)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

var expectedTask = &Task{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/install", Name: "package::install", Environment: struct {
	Name   string "json:\"name,omitempty\""
	CodeID string "json:\"code_id,omitempty\""
}{Name: "production", CodeID: "urn:puppet:code-id:1:a86da166c30f871823f9b2ea224796e834840676;production"}, Metadata: struct {
	Description     string                 "json:\"description,omitempty\""
	SupportsNoop    bool                   "json:\"supports_noop,omitempty\""
	InputMethod     string                 "json:\"input_method,omitempty\""
	Parameters      map[string]TaskParam   "json:\"parameters,omitempty\""
	Extensions      map[string]interface{} "json:\"extensions,omitempty\""
	Implementations []TaskImplementation   "json:\"implementations\""
}{Description: "Bootstrap a node with puppet-agent", SupportsNoop: false, InputMethod: "stdin", Parameters: map[string]TaskParam{"cacert_content": TaskParam{Description: "The expected CA certificate content for the master", Type: "Optional[String]"}, "certname": TaskParam{Description: "The certname with which the node should be bootstrapped", Type: "Optional[String]"}, "custom_attribute": TaskParam{Description: "This setting is added to puppet.conf and included in the custom_attributes section of csr_attributes.yaml", Type: "Optional[Array[Pattern[/\\w+=\\w+/]]]"}, "dns_alt_names": TaskParam{Description: "The DNS alt names with which the agent certificate should be generated", Type: "Optional[String]"}, "environment": TaskParam{Description: "The environment in which the node should be bootstrapped", Type: "Optional[String]"}, "extension_request": TaskParam{Description: "This setting is added to puppet.conf and included in the extension_requests section of csr_attributes.yaml", Type: "Optional[Array[Pattern[/\\w+=\\w+/]]]"}, "master": TaskParam{Description: "The fqdn of the master from which the puppet-agent should be bootstrapped", Type: "String"}, "set_noop": TaskParam{Description: "The noop setting in the [agent] section of puppet.conf", Type: "Optional[Boolean]"}}, Extensions: map[string]interface{}{"discovery": map[string]interface{}{"friendlyName": "Install Puppet agent", "parameters": map[string]interface{}{"cacert_content": map[string]interface{}{"placeholder": "-----BEGIN CERTIFICATE---- ... -----END CERTIFICATE-----"}, "master": map[string]interface{}{"placeholder": "master.company.net"}}, "puppetInstall": true, "type": []interface{}{"host"}}}, Implementations: []TaskImplementation{TaskImplementation{Name: "windows.ps1", Requirements: []string{"powershell"}, InputMethod: "powershell"}, TaskImplementation{Name: "linux.sh", Requirements: []string{"shell"}, InputMethod: "environment"}}}, Files: []struct {
	Filename string "json:\"filename,omitempty\""
	URI      struct {
		Path   string "json:\"path,omitempty\""
		Params struct {
			Environment string "json:\"environment,omitempty\""
		} "json:\"params,omitempty\""
	} "json:\"uri,omitempty\""
	Sha256    string "json:\"sha256,omitempty\""
	SizeBytes int    "json:\"size_bytes,omitempty\""
}{struct {
	Filename string "json:\"filename,omitempty\""
	URI      struct {
		Path   string "json:\"path,omitempty\""
		Params struct {
			Environment string "json:\"environment,omitempty\""
		} "json:\"params,omitempty\""
	} "json:\"uri,omitempty\""
	Sha256    string "json:\"sha256,omitempty\""
	SizeBytes int    "json:\"size_bytes,omitempty\""
}{Filename: "install", URI: struct {
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
}{struct {
	ID   string "json:\"id,omitempty\""
	Name string "json:\"name,omitempty\""
}{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/install", Name: "package::install"}, struct {
	ID   string "json:\"id,omitempty\""
	Name string "json:\"name,omitempty\""
}{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/upgrade", Name: "package::upgrade"}, struct {
	ID   string "json:\"id,omitempty\""
	Name string "json:\"name,omitempty\""
}{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/exec/init", Name: "exec"}}}
