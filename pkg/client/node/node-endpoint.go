package node

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/network"
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

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	e := &network.Endpoint{
		EndpointID: eID,
		Node:       n,
	}

	s := output.Spinner()

	sr, err := nxc.DeleteNetworkEndpoint(context.TODO(), e)
	if err != nil {
		status.Error(err, "Unable to delete node network endpoint")
	}

	s.Stop()

	output.Show(sr)
}
