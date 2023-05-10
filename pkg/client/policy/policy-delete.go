package policy

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/subnet"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	s := subnet.GetSubnet(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ss := output.Spinner()

	sr := &topology.SubnetReq{
		AccountID: s.AccountID,
		TenantID:  s.TenantID,
		NetID:     s.NetID,
		SubnetID:  s.SubnetID,
	}

	np, err := nxc.DeleteNetworkPolicy(context.TODO(), sr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to delete network policy")
	}

	s.NetworkPolicy = np

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
