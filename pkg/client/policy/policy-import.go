package policy

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/ops/object"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils"
)

func (api *API) Import(yamlFile string) {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	npc := &object.NetworkPolicyConfig{}

	if err := utils.FileParser(yamlFile, npc); err != nil {
		status.Error(err, "Unable to parse YAML file")
	}

	s := output.Spinner()

	npcr := &object.NetworkPolicyConfigRequest{
		AccountID:           a.AccountID,
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
