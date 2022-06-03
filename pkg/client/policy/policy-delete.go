package policy

import (
	"context"

	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	v := vrf.GetVRF(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	np, err := nxc.DeleteNetworkPolicy(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete network policy")
	}

	s.Stop()

	// output.Show(np)
	Output().Show(v.TenantID, v.NetID, v.VRFID, np)
}
