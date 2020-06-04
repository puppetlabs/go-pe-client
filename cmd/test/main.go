package main

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/puppetlabs/go-pe-client/pkg/orch"
	"github.com/puppetlabs/go-pe-client/pkg/puppetdb"
	"github.com/puppetlabs/go-pe-client/pkg/rbac"
)

func tokenGesture(peServer string, login string, password string) {

	rbacHostURL := "https://" + peServer + ":4433"
	rbacClient := rbac.NewClient(rbacHostURL, &tls.Config{InsecureSkipVerify: true}) // #nosec - this main() is private and for development purpose
	fmt.Println("Connecting to:", peServer)
	request := &rbac.RequestKeys{
		Login:    login,
		Password: password,
		Label:    "go-pe-client",
	}

	token, err := rbacClient.GetRBACToken(request)
	if err != nil {
		panic(err)
	}
	spew.Dump(token)
	fmt.Println()
}

func main() {

	if len(os.Args) < 3 {
		msg := `Runs through a gamut of PDB and Orchestration queries or returns an RBAC token
Run the queries: go run cmd/main.go <pe-server> <token> e.g. go run cmd/main.go pe.puppetlabs.net aabbccddeeff
or
Get the RBAC token: go run cmd/main.go <pe-server> <login> <password> e.g. go run cmd/main.go pe.puppetlabs.net admin pazzw0rd`
		panic(msg)
	}

	if len(os.Args) == 4 {
		peServer := os.Args[1]
		login := os.Args[2]
		password := os.Args[3]
		tokenGesture(peServer, login, password)
		os.Exit(0)
	}
	peServer := os.Args[1]
	token := os.Args[2]
	pdbHostURL := "https://" + peServer + ":8081"
	pdbClient := puppetdb.NewClient(pdbHostURL, token, &tls.Config{InsecureSkipVerify: true}) // #nosec - this main() is private and for development purpose
	orchHostURL := "https://" + peServer + ":8143"
	orchClient := orch.NewClient(orchHostURL, token, &tls.Config{InsecureSkipVerify: true}) // #nosec - this main() is private and for development purpose
	fmt.Println("Connecting to:", peServer)

	nodes, err := pdbClient.Nodes("", nil, nil)
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	nodes, err = pdbClient.Nodes(fmt.Sprintf(`["=", "certname", "%s"]`, peServer), nil, nil)
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)
	fmt.Println()

	inv, err := orchClient.Inventory()
	if err != nil {
		panic(err)
	}
	spew.Dump(inv)
	fmt.Println()

	jobID, err := orchClient.CommandTask(&orch.TaskRequest{
		Task: "package",
		Params: map[string]string{
			"action": "status",
			"name":   "openssl",
		},
		Scope: orch.Scope{
			Nodes: []string{peServer},
		},
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(jobID)
	fmt.Println()

	stopJobID, err := orchClient.CommandStop(&orch.StopRequest{
		Job: jobID.Job.Name, // Stops the previous task
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(stopJobID)
	fmt.Println()

	scheduledJobID, err := orchClient.CommandScheduleTask(&orch.ScheduleTaskRequest{
		Task: "package",
		Params: map[string]string{
			"action": "status",
			"name":   "openssl",
		},
		Scope: orch.Scope{
			Nodes: []string{peServer},
		},
		ScheduledTime: "2027-05-05T19:50:08Z",
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(scheduledJobID)
	fmt.Println()

	taskTargetJobID, err := orchClient.CommandTaskTarget(&orch.TaskTargetRequest{
		DisplayName: "foo",
		AllTasks:    true,
		Nodes:       []string{"node1", "node2", "node3"},
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(taskTargetJobID)
	fmt.Println()

	planRunJobID, err := orchClient.CommandPlanRun(&orch.PlanRunRequest{
		Name:        "foo",
		Environment: "production",
		Description: "Optional description",
		Params: map[string]interface{}{
			"string": "test",
			"number": 111,
			"list":   []string{"one", "two", "three"},
		},
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(planRunJobID)
	fmt.Println()

	JobID, err := orchClient.CommandDeploy(&orch.DeployRequest{
		Environment: "production",
		Noop:        true,
		NoNoop:      false,
		Scope: orch.Scope{
			Nodes: []string{"node1.example.com"},
		},
		Concurrency:        2,
		Description:        "Description of this job",
		EnforceEnvironment: true,
		Trace:              true,
		Evaltrace:          false,
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(JobID)
	fmt.Println()

	job, err := orchClient.Job(jobID.Job.Name)
	if err != nil {
		panic(err)
	}
	spew.Dump(job)
	fmt.Println()

	jobReport, err := orchClient.JobReport(jobID.Job.Name)
	if err != nil {
		panic(err)
	}
	spew.Dump(jobReport)
	fmt.Println()

	jobNodes, err := orchClient.JobNodes(jobID.Job.Name)
	if err != nil {
		panic(err)
	}
	spew.Dump(jobNodes)
	fmt.Println()

	tasks, err := orchClient.Tasks("production")
	if err != nil {
		panic(err)
	}
	spew.Dump(tasks)

	plans, err := orchClient.Plans("production")
	if err != nil {
		panic(err)
	}
	spew.Dump(plans)
}
