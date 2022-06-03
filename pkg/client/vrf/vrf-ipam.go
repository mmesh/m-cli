package vrf

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

type ipamEndpoint struct {
	endpointID string
	ipv4       string
}

func (api *API) DeleteIPAMEntry() {
	v := GetVRF(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	// endpt := make(map[string]string) // map[endpointID]ipv4

	var endpointOptID string
	// var ipamEndpoints []string

	ipamEndpointsOpts := make([]string, 0)
	ipamEndpoints := make(map[string]*ipamEndpoint)

	for e, ipv4 := range v.IPAM.Endpoints {
		endpointOptID = fmt.Sprintf("%s: %s", ipv4, e)
		ipamEndpointsOpts = append(ipamEndpointsOpts, endpointOptID)
		ipamEndpoints[endpointOptID] = &ipamEndpoint{
			endpointID: e,
			ipv4:       ipv4,
		}
		// ipamEndpoints = append(ipamEndpoints, e)
		// endpt[e] = ipv4
	}

	if len(ipamEndpointsOpts) == 0 {
		msg.Info("No objects found")
		return
	}

	endpointOptID = input.GetSelect("IPAM Endpoint:", "", ipamEndpointsOpts, survey.Required)

	s := output.Spinner()

	// delete(v.IPAM.Endpoints, eID)
	delete(v.IPAM.Endpoints, ipamEndpoints[endpointOptID].endpointID)
	// delete(v.IPAM.Leased, endpt[eID])
	delete(v.IPAM.Leased, ipamEndpoints[endpointOptID].ipv4)
	// v.IPAM.Available[endpt[eID]] = true
	v.IPAM.Available[ipamEndpoints[endpointOptID].ipv4] = true

	vrf, err := nxc.SetVRF(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set VRF (subnet)")
	}

	s.Stop()

	// output.Show(v)
	Output().Show(vrf)
}
