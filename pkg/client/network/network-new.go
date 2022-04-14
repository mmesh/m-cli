package network

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/tenant"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) New() {
	t := tenant.GetTenant(false)

	n := &network.Network{
		AccountID: t.AccountID,
		TenantID:  t.TenantID,
	}

	n.NetID = input.GetInput("Network ID:", "", "", validNetID)

	helpMsg := "A valid /16 network with format 'n.n.0.0/16' is required"
	n.NetworkCIDR = input.GetInput("Network CIDR:", helpMsg, "", input.ValidNetwork)

	n.Description = input.GetInput("Description:", "", n.Description, survey.Required)

	n.RouteVRFs = input.GetConfirm("Route this network's subnets (VRFs) each other?", n.RouteVRFs)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	n, err := nxc.SetNetwork(context.TODO(), n)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create network")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(n)
}
