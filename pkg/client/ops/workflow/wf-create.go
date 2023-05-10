package workflow

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/client/ops/project"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils"
)

func (api *API) Create(yamlFile string) {
	p := project.GetProject()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	wf := &ops.Workflow{}

	if err := utils.FileParser(yamlFile, wf); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	wf.AccountID = p.AccountID
	wf.TenantID = p.TenantID
	wf.ProjectID = p.ProjectID

	wf, err := nxc.CreateWorkflow(context.TODO(), wf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create workflow")
	}

	s.Stop()

	Output().Show(wf)
}
