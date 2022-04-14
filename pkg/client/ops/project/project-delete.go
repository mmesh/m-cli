package project

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	p := GetProject(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sr, err := nxc.DeleteProject(context.TODO(), p)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete project")
	}

	s.Stop()

	output.Show(sr)
}
