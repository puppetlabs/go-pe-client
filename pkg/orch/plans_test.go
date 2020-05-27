package orch

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/require"
)

func TestPlans(t *testing.T) {

	// Test without environment
	setupGetResponder(t, orchPlans, "", "plans-response.json")
	actual, err := orchClient.Plans("")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test with environment
	setupGetResponder(t, orchPlans, "environment=myenv", "plans-response.json")
	actual, err = orchClient.Plans("myenv")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, orchPlans)
	actual, err = orchClient.Plans("")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestPlan(t *testing.T) {

	replacer := strings.NewReplacer("{module}", "package", "{planname}", "install")
	orchPlanPackageInstall := replacer.Replace(orchPlanName)

	// Test without environment
	setupGetResponder(t, orchPlanPackageInstall, "", "plan-response.json")
	actual, err := orchClient.Plan("", "package", "install")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test with environment
	setupGetResponder(t, orchPlanPackageInstall, "environment=myenv", "plan-response.json")
	actual, err = orchClient.Plan("myenv", "package", "install")
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, orchPlanPackageInstall)
	actual, err = orchClient.Plan("myenv", "package", "install")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestPlanByID(t *testing.T) {
	replacer := strings.NewReplacer("{module}", "package", "{planname}", "upgrade")
	orchPlanPackageUpgrade := replacer.Replace(orchPlanName)

	id := "https://orchestrator.example.com:8143" + orchPlanPackageUpgrade

	// Test without environment
	setupGetResponder(t, orchPlanPackageUpgrade, "", "plan-response.json")
	actual, err := orchClient.PlanByID("", id)
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test with environment
	setupGetResponder(t, orchPlanPackageUpgrade, "environment=myenv", "plan-response.json")
	actual, err = orchClient.PlanByID("myenv", id)
	require.Nil(t, err)
	require.False(t, structs.HasZero(actual), spew.Sdump(actual))

	// Test error
	setupErrorResponder(t, orchPlanPackageUpgrade)
	actual, err = orchClient.PlanByID("myenv", id)
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}
