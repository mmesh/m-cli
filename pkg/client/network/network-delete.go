package network

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	n := GetNetwork(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sr, err := nxc.DeleteNetwork(context.TODO(), n)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete network")
	}

	s.Stop()

	output.Show(sr)
}
