package acl

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Show() {
	a := account.GetAccount()

	nxc, grpcConn1 := grpc.GetTopologyAPIClient()
	defer grpcConn1.Close()

	s := output.Spinner()

	req := &topology.TopologyRequest{
		AccountID: a.AccountID,
	}

	nsm, err := nxc.GetNodeSummaryMap(context.TODO(), req)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get all-nodes summary map")
	}

	s.Stop()

	Output().Show(GetACL(false), nsm)
}
