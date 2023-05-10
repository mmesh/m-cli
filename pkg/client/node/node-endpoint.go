package node

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) ShowEndpoint() {
	n := GetNode(false)
	eID := GetEndpoint(n).EndpointID

	for endpointID, e := range n.Endpoints {
		if endpointID == eID {
			//output.Show(e)
			Output().ShowEndpoint(e)
		}
	}
}

func (api *API) DeleteEndpoint() {
	n := GetNode(false)
	eID := GetEndpoint(n).EndpointID

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	er := &topology.EndpointRequest{
		NodeReq: &topology.NodeReq{
			AccountID: n.AccountID,
			TenantID:  n.TenantID,
			NetID:     n.NetID,
			SubnetID:  n.SubnetID,
			NodeID:    n.NodeID,
		},
		EndpointID: eID,
	}

	sr, err := nxc.DeleteNetworkEndpoint(context.TODO(), er)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete node network endpoint")
	}

	s.Stop()

	output.Show(sr)
}
