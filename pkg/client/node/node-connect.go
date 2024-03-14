package node

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/subnet"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/mm"
)

func (api *API) Connect() {
	n := GetNodeByTenant(false, nil)

	s := subnet.GetSubnet(false)

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNodeNetworkingRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		NetID:    s.NetID,
		SubnetID: s.SubnetID,
	}

	sr, err := nxc.UpdateNodeNetworking(context.TODO(), unr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to update node network configuration")
	}

	ss.Stop()

	// output.Show(sr)
	Output().Show(sr)
}

func (api *API) Disconnect() {
	n := GetNodeByTenant(false, mm.Bool(true))

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	unr := &topology.UpdateNodeNetworkingRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NodeID:    n.NodeID,
		},
		NetID:    "",
		SubnetID: "",
	}

	sr, err := nxc.UpdateNodeNetworking(context.TODO(), unr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update node network configuration")
	}

	s.Stop()

	// output.Show(sr)
	Output().Show(sr)
}
