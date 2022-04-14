package node

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	n := GetNode(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sr, err := nxc.DeleteNode(context.TODO(), n)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete node")
	}

	s.Stop()

	output.Show(sr)
}
