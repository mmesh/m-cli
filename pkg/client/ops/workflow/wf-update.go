package workflow

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils"
)

func (api *API) Update(yamlFile string) {
	wf := GetWorkflow()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	nwf := &ops.Workflow{}

	if err := utils.FileParser(yamlFile, nwf); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	nwf.AccountID = wf.AccountID
	nwf.TenantID = wf.TenantID
	nwf.ProjectID = wf.ProjectID

	if nwf.WorkflowID != wf.WorkflowID {
		s.Stop()
		err := fmt.Errorf("workflowID mismatch")
		status.Error(err, "WorkflowID is not valid")
	}

	wf, err := nxc.SetWorkflow(context.TODO(), nwf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set workflow")
	}

	s.Stop()

	Output().Show(wf)
}
