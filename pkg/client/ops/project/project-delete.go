package project

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	p := GetProject()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	pr := &ops.ProjectReq{
		AccountID: p.AccountID,
		TenantID:  p.TenantID,
		ProjectID: p.ProjectID,
	}

	sr, err := nxc.DeleteProject(context.TODO(), pr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete project")
	}

	s.Stop()

	output.Show(sr)
}
