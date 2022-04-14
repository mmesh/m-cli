package vrf

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	network_pb "mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/network"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/ipnet"
)

func (api *API) New() {
	n := network.GetNetwork(false)

	v := &network_pb.VRF{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
		IPAM: &network_pb.IPAM{
			NetworkCIDR: n.NetworkCIDR,
		},
	}

	// global var needed by the validation function
	networkCIDR = n.NetworkCIDR

	helpText, err := subnetHelp(networkCIDR)
	if err != nil {
		status.Error(err, "Unable to parse network CIDR")
	}

	v.IPAM.SubnetCIDR = input.GetInput("Subnet CIDR:", helpText, "", validSubnet)

	subnetID, err := getSubnetID(v.IPAM.SubnetCIDR)
	if err != nil {
		status.Error(err, "Unable to parse subnet CIDR")
	}

	v.VRFID = subnetID

	v.NetworkPolicy = &network_pb.Policy{
		DefaultPolicy:  ipnet.Policy_ACCEPT,
		NetworkFilters: make([]*network_pb.Filter, 0),
	}

	v.Description = input.GetInput("Subnet Description:", "", v.Description, survey.Required)

	pwMsg := "This secret will never be shown, keep it safe"
	v.AuthSecret = input.GetPassword("Subnet Secret:", pwMsg)

	if v.NetworkPolicy == nil {
		v.NetworkPolicy = &network_pb.Policy{
			DefaultPolicy:  ipnet.Policy_ACCEPT,
			NetworkFilters: make([]*network_pb.Filter, 0),
		}
	}
	np := v.NetworkPolicy
	np.DefaultPolicy = input.GetSelect("Default Network Policy:", "", vars.Policies, survey.Required)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	v, err = nxc.SetVRF(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create subnet")
	}

	s.Stop()

	// output.Show(v)
	Output().Show(v)
}
