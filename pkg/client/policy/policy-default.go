package policy

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/ops/object"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) SetDefault() {
	v := vrf.GetVRF(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	npc := setDefaultNetworkPolicy(v)

	s := output.Spinner()

	npcr := &object.NetworkPolicyConfigRequest{
		AccountID:           v.AccountID,
		NetworkPolicyConfig: npc,
	}

	np, err := nxc.SetNetworkPolicy(context.TODO(), npcr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set network policy")
	}

	s.Stop()

	// output.Show(np)
	Output().Show(np)
}
