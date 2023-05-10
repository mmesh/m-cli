package subnet

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/topology"
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
	s := GetSubnet(false)

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	var endpointOptID string

	ipamEndpointsOpts := make([]string, 0)
	ipamEndpoints := make(map[string]*ipamEndpoint)

	for e, ipv4 := range s.IPAM.Endpoints {
		endpointOptID = fmt.Sprintf("%s: %s", ipv4, e)
		ipamEndpointsOpts = append(ipamEndpointsOpts, endpointOptID)
		ipamEndpoints[endpointOptID] = &ipamEndpoint{
			endpointID: e,
			ipv4:       ipv4,
		}
	}

	if len(ipamEndpointsOpts) == 0 {
		msg.Info("No objects found")
		return
	}

	endpointOptID = input.GetSelect("IPAM Endpoint:", "", ipamEndpointsOpts, survey.Required)

	ss := output.Spinner()

	req := &topology.RemoveSubnetIPAMEntryRequest{
		AccountID:  s.AccountID,
		TenantID:   s.TenantID,
		NetID:      s.NetID,
		SubnetID:   s.SubnetID,
		EndpointID: ipamEndpoints[endpointOptID].endpointID,
		IPv4:       ipamEndpoints[endpointOptID].ipv4,
	}

	s, err := nxc.RemoveSubnetIPAMEntry(context.TODO(), req)
	if err != nil {
		ss.Stop()
		status.Error(err, "Unable to remove subnet IPAM entry")
	}

	ss.Stop()

	// output.Show(s)
	Output().Show(s)
}
