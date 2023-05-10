package node

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	n := GetNode(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	nr := &topology.NodeReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
		SubnetID:  n.SubnetID,
		NodeID:    n.NodeID,
	}

	sr, err := nxc.DeleteNode(context.TODO(), nr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete node")
	}

	s.Stop()

	output.Show(sr)
}
