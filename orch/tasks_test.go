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

var expectedTask = &Task{ID: "https://orchestrator.example.com:8143/orchestrator/v1/tasks/package/install", Name: "package::install", Environment: struct {
	Name   string "json:\"name,omitempty\""
	CodeID string "json:\"code_id,omitempty\""
}{Name: "production", CodeID: "urn:puppet:code-id:1:a86da166c30f871823f9b2ea224796e834840676;production"}, Metadata: struct {
	Description  string "json:\"description,omitempty\""
	SupportsNoop bool   "json:\"supports_noop,omitempty\""
	InputMethod  string "json:\"input_method,omitempty\""
	Parameters   struct {
		Name struct {
			Description string "json:\"description,omitempty\""
			Type        string "json:\"type,omitempty\""
		} "json:\"name,omitempty\""
		Provider struct {
			Description string "json:\"description,omitempty\""
			Type        string "json:\"type,omitempty\""
		} "json:\"provider,omitempty\""
		Version struct {
			Description string "json:\"description,omitempty\""
			Type        string "json:\"type,omitempty\""
		} "json:\"version,omitempty\""
	} "json:\"parameters,omitempty\""
}{Description: "Install a package", SupportsNoop: true, InputMethod: "stdin", Parameters: struct {
	Name struct {
		Description string "json:\"description,omitempty\""
		Type        string "json:\"type,omitempty\""
	} "json:\"name,omitempty\""
	Provider struct {
		Description string "json:\"description,omitempty\""
		Type        string "json:\"type,omitempty\""
	} "json:\"provider,omitempty\""
	Version struct {
		Description string "json:\"description,omitempty\""
		Type        string "json:\"type,omitempty\""
	} "json:\"version,omitempty\""
}{Name: struct {
	Description string "json:\"description,omitempty\""
	Type        string "json:\"type,omitempty\""
}{Description: "The package to install", Type: "String[1]"}, Provider: struct {
	Description string "json:\"description,omitempty\""
	Type        string "json:\"type,omitempty\""
}{Description: "The provider to use to install the package", Type: "Optional[String[1]]"}, Version: struct {
	Description string "json:\"description,omitempty\""
	Type        string "json:\"type,omitempty\""
}{Description: "The version of the package to install, defaults to latest", Type: "Optional[String[1]]"}}}, Files: []struct {
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
