package network

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Update() {
	n := GetNetwork(false)

	n.Description = input.GetInput("Description:", "", n.Description, survey.Required)

	n.RouteVRFs = input.GetConfirm("Route this network's subnets (VRFs) each other?", n.RouteVRFs)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	n, err := nxc.SetNetwork(context.TODO(), n)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update network")
	}

	s.Stop()

	// output.Show(n)
	Output().Show(n)
}
