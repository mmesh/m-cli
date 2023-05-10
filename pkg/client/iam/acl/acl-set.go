package acl

import (
	"context"
	"fmt"
	"sort"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	nxc1, grpcConn1 := grpc.GetTopologyAPIClient()
	defer grpcConn1.Close()

	acl := GetACL(true)
	if acl != nil { // editing existing resource
		output.Choice("Edit RBAC ACL")
	} else { // <new> resource
		output.Choice("New RBAC ACL")

		acl = &iam.ACL{
			AccountID: a.AccountID,
			Users:     make(map[string]bool),
		}

		acl.ACLID = input.GetInput("ACL ID:", "", "", validACLID)
	}

	s := output.Spinner()

	req := &topology.TopologyRequest{
		AccountID: a.AccountID,
	}

	nsm, err := nxc1.GetNodeSummaryMap(context.TODO(), req)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get all-nodes summary")
	}

	// get all-node options
	nodesOpts := make([]string, 0)
	nodes := make(map[string]*topology.NodeSummary)

	for _, n := range nsm.Nodes {
		nodeOptID := fmt.Sprintf("[%s] %s: %s", n.TenantName, n.SubnetID, n.NodeName)
		nodesOpts = append(nodesOpts, nodeOptID)
		nodes[nodeOptID] = n
	}

	sort.Strings(nodesOpts)

	// get current acl nodes
	currentNodesOpts := make([]string, 0)
	for _, nodeID := range acl.NodeIDs {
		if n, ok := nsm.Nodes[nodeID]; ok {
			nodeOptID := fmt.Sprintf("[%s] %s: %s", n.TenantName, n.SubnetID, n.NodeName)
			currentNodesOpts = append(currentNodesOpts, nodeOptID)
		}
	}

	s.Stop()

	nodesOpts = input.GetMultiSelect("Nodes:", "", nodesOpts, currentNodesOpts, nil)

	nodeIDs := make([]string, 0)

	for _, nodeOptID := range nodesOpts {
		if ns, ok := nodes[nodeOptID]; ok {
			nodeIDs = append(nodeIDs, ns.NodeID)
		}
	}

	acl.NodeIDs = nodeIDs

	s = output.Spinner()

	nxc2, grpcConn2 := grpc.GetIAMAPIClient()
	defer grpcConn2.Close()

	acl, err = nxc2.SetACL(context.TODO(), acl)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set ACL")
	}

	s.Stop()

	Output().Show(acl, nsm)
}
