package vrf

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) DeleteIPAMEntry() {
	v := GetVRF(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	endpt := make(map[string]string) // map[endpointID]ipv4

	var eID string
	var ipamEndpoints []string
	for e, ipv4 := range v.IPAM.Endpoints {
		ipamEndpoints = append(ipamEndpoints, e)
		endpt[e] = ipv4
	}

	eID = input.GetSelect("IPAM Endpoint:", "", ipamEndpoints, survey.Required)

	s := output.Spinner()

	delete(v.IPAM.Endpoints, eID)
	delete(v.IPAM.Leased, endpt[eID])
	v.IPAM.Available[endpt[eID]] = true

	vrf, err := nxc.SetVRF(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set VRF (subnet)")
	}

	s.Stop()

	// output.Show(v)
	Output().Show(vrf)
}
