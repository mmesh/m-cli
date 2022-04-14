package operation

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	o := GetOperation()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteOperation(context.TODO(), o)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete operation")
	}

	s.Stop()

	output.Show(sr)
}
