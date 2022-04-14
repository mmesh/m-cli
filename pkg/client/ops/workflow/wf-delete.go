package workflow

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	wf := GetWorkflow(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteWorkflow(context.TODO(), wf)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete workflow")
	}

	s.Stop()

	output.Show(sr)
}
