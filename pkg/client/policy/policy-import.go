package policy

/*
import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils"
)

func (api *API) Import(yamlFile string) {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	s := output.Spinner()

	snpr := &topology.SetNetworkPolicyRequest{}

	if err := utils.FileParser(yamlFile, snpr); err != nil {
		s.Stop()
		status.Error(err, "Unable to parse YAML file")
	}

	snpr.AccountID = a.AccountID

	np, err := nxc.SetNetworkPolicy(context.TODO(), snpr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set network policy")
	}

	s.Stop()

	// output.Show(np)
	Output().Show(snpr.TenantID, snpr.NetID, snpr.SubnetID, np)
}
*/
