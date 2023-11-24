package node

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/subnet"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetNode(edit bool) *topology.Node {
	nl := nodes()

	if len(nl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var nodeOptID string
	nodesOpts := make([]string, 0)
	nodes := make(map[string]*topology.Node)

	for nodeName, n := range nl {
		// nodeOptID = nodeName
		nodeOptID = fmt.Sprintf("[%s] %s", nodeName, n.Cfg.Description)
		nodesOpts = append(nodesOpts, nodeOptID)
		nodes[nodeOptID] = n
	}

	sort.Strings(nodesOpts)

	if edit {
		nodesOpts = append(nodesOpts, input.NewResource)
	}

	nodeOptID = input.GetSelect("Node:", "", nodesOpts, survey.Required)

	if nodeOptID == input.NewResource {
		return nil
	}

	vars.NodeID = nodes[nodeOptID].NodeID

	return nodes[nodeOptID]
}

func GetEndpoint(n *topology.Node) *topology.Endpoint {
	var eID string
	var endpoints []string

	for endpointID := range n.Endpoints {
		endpoints = append(endpoints, endpointID)
	}

	sort.Strings(endpoints)

	if len(endpoints) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	eID = input.GetSelect("Endpoints:", "", endpoints, survey.Required)

	return n.Endpoints[eID]
}

func nodes() map[string]*topology.Node {
	s := subnet.GetSubnet(false)

	ss := output.Spinner()
	defer ss.Stop()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	lr := &topology.ListNodesBySubnetRequest{
		Meta: &resource.ListRequest{},
		Subnet: &topology.SubnetReq{
			AccountID: s.AccountID,
			TenantID:  s.TenantID,
			NetID:     s.NetID,
			SubnetID:  s.SubnetID,
		},
	}

	nodes := make(map[string]*topology.Node) // map[nodeName]*topology.Node

	for {
		nl, err := nxc.ListNodesBySubnet(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list nodes by subnet")
		}

		for _, n := range nl.Nodes {
			if n.Cfg != nil {
				if len(n.Cfg.NodeName) > 0 {
					nodes[n.Cfg.NodeName] = n
				}
			}
		}

		if len(nl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = nl.Meta.NextPageToken
		} else {
			break
		}
	}

	return nodes
}

func FetchNode(nr *topology.NodeReq) *topology.Node {
	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	n, err := nxc.GetNode(context.TODO(), nr)
	if err != nil {
		status.Error(err, "Unable to get node")
	}

	return n
}
