package vrf

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	v := GetVRF(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sr, err := nxc.DeleteVRF(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete VRF (subnet)")
	}

	s.Stop()

	output.Show(sr)
}
