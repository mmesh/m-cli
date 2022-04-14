package workflow

import (
	"context"
	"errors"

	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-cli/pkg/client/ops/project"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils"
)

func (api *API) Set(yamlFile string) {
	p := project.GetProject(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	wf := &workflow.Workflow{}

	if err := utils.FileParser(yamlFile, wf); err != nil {
		status.Error(err, "Unable to parse YAML file")
	}

	if p.ProjectID != wf.Project {
		err := errors.New("projectID does not match the one in file")
		status.Error(err, "Invalid ProjectID")
	}

	if err := validWorkflowID(wf.WorkflowID); err != nil {
		status.Error(err, "Invalid WorkflowID")
	}

	wf.AccountID = p.AccountID
	wf.ProjectID = p.ProjectID

	s := output.Spinner()

	wf, err := nxc.SetWorkflow(context.TODO(), wf)
	if err != nil {
		status.Error(err, "Unable to set workflow")
	}

	s.Stop()

	Output().Show(wf)
}
