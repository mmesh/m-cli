package subnet

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	s := GetSubnet(false)

	output.ConfirmDeletion()

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	sr := &topology.SubnetReq{
		AccountID: s.AccountID,
		TenantID:  s.TenantID,
		NetID:     s.NetID,
		SubnetID:  s.SubnetID,
	}

	r, err := nxc.DeleteSubnet(context.TODO(), sr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to delete subnet")
	}

	ss.Stop()

	output.Show(r)
}
