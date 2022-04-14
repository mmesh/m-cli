package cluster

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) DeleteNode() {
	kc := GetKubernetesCluster(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	// select nodepool
	nodePools := make([]string, 0)

	for npID := range kc.Spec.NodePools {
		nodePools = append(nodePools, npID)
	}

	sort.Strings(nodePools)

	if len(nodePools) == 0 {
		msg.Alert("No object found")
		os.Exit(1)
	}

	nodePoolID := input.GetSelect("Node Pool ID:", "", nodePools, survey.Required)

	// select node
	nodes := make([]string, 0)

	for nID := range kc.Spec.NodePools[nodePoolID].Nodes {
		nodes = append(nodes, nID)
	}

	sort.Strings(nodes)

	if len(nodes) == 0 {
		msg.Alert("No object found")
		os.Exit(1)
	}

	if len(nodePools) == 1 && len(nodes) == 1 {
		msg.Info("Cannot delete the last worker node in cluster")
		os.Exit(0)
	}

	nodeID := input.GetSelect("Node ID:", "", nodes, survey.Required)

	// build request
	kndReq := &cloud.KubernetesNodeDeleteRequest{
		AccountID:           kc.AccountID,
		KubernetesClusterID: kc.KubernetesClusterID,
		ServiceID:           kc.ServiceID,
		ProviderID:          kc.ProviderID,
		NodePool: &cloud.KubernetesNodePool{
			NodePoolID: nodePoolID,
		},
		Node: &cloud.KubernetesNode{
			KubernetesNodeID: nodeID,
		},
	}

	output.ConfirmDeletion()

	s := output.Spinner()

	kc, err := nxc.DeleteCloudKubernetesNode(context.TODO(), kndReq)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete kubernetes node pool")
	}

	s.Stop()

	Output().Show(kc)
}
