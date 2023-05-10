package network

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	n := GetNetwork(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	nr := &topology.NetworkReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
	}

	sr, err := nxc.DeleteNetwork(context.TODO(), nr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete network")
	}

	s.Stop()

	output.Show(sr)
}
