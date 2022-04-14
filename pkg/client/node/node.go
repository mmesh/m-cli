package node

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetNode(edit bool) *network.Node {
	nl := nodes()

	if len(nl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var nodeOptID string
	nodesOpts := make([]string, 0)
	nodes := make(map[string]*network.Node)

	for _, n := range nl {
		nodeOptID = n.NodeID
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

func GetEndpoint(n *network.Node) *network.NetworkEndpoint {
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

func nodes() map[string]*network.Node {
	v := vrf.GetVRF(false)

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &network.ListNodesRequest{
		Meta: &resource.ListRequest{},
		VRF:  v,
	}

	nodes := make(map[string]*network.Node)

	for {
		nl, err := nxc.ListNodes(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list nodes")
		}

		for _, n := range nl.Nodes {
			nodes[n.NodeID] = n
		}

		if len(nl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = nl.Meta.NextPageToken
		} else {
			break
		}
	}

	return nodes
}
