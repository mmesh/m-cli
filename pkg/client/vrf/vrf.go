package vrf

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	network_pb "mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/network"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var networkCIDR string
var vrfsMap map[string]*network_pb.VRF = nil

func GetVRF(edit bool) *network_pb.VRF {
	vl := vrfs()

	if len(vl) == 0 {
		msg.Info("No subnets found")
		os.Exit(1)
	}

	var vrfOptID string
	vrfsOpts := make([]string, 0)
	vrfs := make(map[string]*network_pb.VRF)

	for _, v := range vl {
		if v.IPAM != nil {
			vrfOptID = fmt.Sprintf("[%s] %s", v.VRFID, v.IPAM.SubnetCIDR)
		} else {
			vrfOptID = v.VRFID
		}
		vrfsOpts = append(vrfsOpts, vrfOptID)
		vrfs[vrfOptID] = v
	}

	sort.Strings(vrfsOpts)

	if edit {
		vrfsOpts = append(vrfsOpts, input.NewResource)
	}

	vrfOptID = input.GetSelect("Subnet:", "", vrfsOpts, survey.Required)

	if vrfOptID == input.NewResource {
		return nil
	}

	vars.VRFID = vrfs[vrfOptID].VRFID

	return vrfs[vrfOptID]
}

func vrfs() map[string]*network_pb.VRF {
	n := network.GetNetwork(false)

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &network_pb.ListVRFsRequest{
		Meta:    &resource.ListRequest{},
		Network: n,
	}

	vrfs := make(map[string]*network_pb.VRF)

	for {
		vl, err := nxc.ListVRFs(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list subnets")
		}

		for _, v := range vl.VRFs {
			vrfs[v.VRFID] = v
		}

		if len(vl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = vl.Meta.NextPageToken
		} else {
			break
		}
	}

	return vrfs
}

func validSubnet(val interface{}) error {
	_, ipv4Network, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		return err
	}

	ipv4Addr, ipv4Subnet, err := net.ParseCIDR(val.(string))
	if err != nil {
		return err
	}

	if !ipv4Addr.Equal(ipv4Subnet.IP) {
		return fmt.Errorf("invalid subnet address %v", ipv4Addr)
	}

	if !ipv4Network.Contains(ipv4Subnet.IP) {
		return fmt.Errorf("subnet %s is not included in network %s", val.(string), networkCIDR)
	}

	cidrMask, _ := ipv4Subnet.Mask.Size()

	if cidrMask != 24 {
		return errors.New("only /24 subnets are supported at the moment")
	}

	subnetCIDR := val.(string)

	subnetID, err := getSubnetID(subnetCIDR)
	if err != nil {
		return fmt.Errorf("unable to parse subnet CIDR: %v", err)
	}

	if vrfsMap == nil {
		vrfsMap = vrfs()
	}

	if _, ok := vrfsMap[subnetID]; ok {
		return fmt.Errorf("subnet %s already exist", subnetCIDR)
	}

	return nil
}

func getSubnetID(subnetCIDR string) (string, error) {
	_, ipv4Net, err := net.ParseCIDR(subnetCIDR)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("subnet-%03d", ipv4Net.IP[2]), nil
}

func subnetHelp(networkCIDR string) (string, error) {
	_, ipv4Net, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%d.%d.n.0/24", ipv4Net.IP[0], ipv4Net.IP[1])

	text := fmt.Sprintf("Network %s: A valid /24 subnet with format '%s' is required", networkCIDR, s)

	return text, nil
}
