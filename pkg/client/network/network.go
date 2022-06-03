package network

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/tenant"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var networksMap map[string]*network.Network = nil
var selectedNetwork *network.Network = nil

func GetNetwork(edit bool) *network.Network {
	if selectedNetwork != nil {
		return selectedNetwork
	}

	nl := networks()

	if len(nl) == 0 {
		msg.Info("No networks found")
		os.Exit(1)
	}

	var networkOptID string
	networksOpts := make([]string, 0)
	networks := make(map[string]*network.Network)

	for _, n := range nl {
		networkOptID = fmt.Sprintf("[%s] %s", n.NetID, n.NetworkCIDR)
		networksOpts = append(networksOpts, networkOptID)
		networks[networkOptID] = n
	}

	sort.Strings(networksOpts)

	if edit {
		networksOpts = append(networksOpts, input.NewResource)
	}

	networkOptID = input.GetSelect("Network:", "", networksOpts, survey.Required)

	if networkOptID == input.NewResource {
		return nil
	}

	vars.NetID = networks[networkOptID].NetID
	selectedNetwork = networks[networkOptID]

	return networks[networkOptID]
}

func networks() map[string]*network.Network {
	if networksMap != nil {
		return networksMap
	}

	t := tenant.GetTenant(false)

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &network.ListNetworksRequest{
		Meta:   &resource.ListRequest{},
		Tenant: t,
	}

	networks := make(map[string]*network.Network)

	for {
		nl, err := nxc.ListNetworks(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list networks")
		}

		for _, n := range nl.Networks {
			networks[n.NetID] = n
		}

		if len(nl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = nl.Meta.NextPageToken
		} else {
			break
		}
	}

	networksMap = networks

	return networks
}

func validNetID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	netID := val.(string)

	if networksMap == nil {
		networksMap = networks()
	}

	if _, ok := networksMap[netID]; ok {
		return fmt.Errorf("network %s already exist", netID)
	}

	return nil
}
