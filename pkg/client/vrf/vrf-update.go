package vrf

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/ipnet"
)

func (api *API) Update() {
	v := GetVRF(false)

	v.Description = input.GetInput("Subnet Description:", "", v.Description, survey.Required)

	pwMsg := "This secret will never be shown, keep it safe"
	v.AuthSecret = input.GetPassword("Subnet Secret:", pwMsg)

	if v.NetworkPolicy == nil {
		v.NetworkPolicy = &network.Policy{
			DefaultPolicy:  ipnet.Policy_ACCEPT,
			NetworkFilters: make([]*network.Filter, 0),
		}
	}
	np := v.NetworkPolicy
	np.DefaultPolicy = input.GetSelect("Default Network Policy:", "", vars.Policies, survey.Required)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	v, err := nxc.SetVRF(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to update subnet")
	}

	s.Stop()

	// output.Show(v)
	Output().Show(v)
}
