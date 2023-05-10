package subnet

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/network"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) New() {
	n := network.GetNetwork(false)

	nsr := &topology.NewSubnetRequest{
		AccountID:   n.AccountID,
		TenantID:    n.TenantID,
		NetID:       n.NetID,
		NetworkCIDR: n.NetworkCIDR,
	}

	// global var needed by the validation function
	networkCIDR = n.NetworkCIDR

	helpText, err := subnetHelp(networkCIDR)
	if err != nil {
		status.Error(err, "Unable to parse network CIDR")
	}

	nsr.SubnetCIDR = input.GetInput("Subnet CIDR:", helpText, "", validSubnet)

	nsr.Description = input.GetInput("Subnet Description:", "", "", survey.Required)

	nsr.DefaultPolicy = GetSecurityPolicy("Default Security Policy:")

	ss := output.Spinner()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	s, err := nxc.CreateSubnet(context.TODO(), nsr)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to create subnet")
	}

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
