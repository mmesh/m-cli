package workflow

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/ops/project"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetWorkflow(edit bool) *workflow.Workflow {
	wfl := workflows()

	if len(wfl) == 0 {
		msg.Info("No workflows found")
		os.Exit(1)
	}

	var workflowOptID string
	workflowsOpts := make([]string, 0)
	workflows := make(map[string]*workflow.Workflow)

	for _, wf := range wfl {
		workflowOptID = wf.WorkflowID
		workflowsOpts = append(workflowsOpts, workflowOptID)
		workflows[workflowOptID] = wf
	}

	sort.Strings(workflowsOpts)

	if edit {
		workflowsOpts = append(workflowsOpts, input.NewResource)
	}

	workflowOptID = input.GetSelect("Workflow:", "", workflowsOpts, survey.Required)

	if workflowOptID == input.NewResource {
		return nil
	}

	vars.WorkflowID = workflows[workflowOptID].WorkflowID

	return workflows[workflowOptID]
}

func workflows() map[string]*workflow.Workflow {
	p := project.GetProject(false)

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &workflow.ListWorkflowsRequest{
		Meta:    &resource.ListRequest{},
		Project: p,
	}

	workflows := make(map[string]*workflow.Workflow)

	for {
		wfl, err := nxc.ListWorkflows(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list workflows")
		}

		for _, wf := range wfl.Workflows {
			workflows[wf.WorkflowID] = wf
		}

		if len(wfl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = wfl.Meta.NextPageToken
		} else {
			break
		}
	}

	return workflows
}

func validWorkflowID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	return nil
}
