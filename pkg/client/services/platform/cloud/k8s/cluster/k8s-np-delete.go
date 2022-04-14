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

func (api *API) DeleteNodePool() {
	kc := GetKubernetesCluster(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	nodePools := make([]string, 0)

	for npID := range kc.Spec.NodePools {
		nodePools = append(nodePools, npID)
	}

	sort.Strings(nodePools)

	if len(nodePools) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	if len(nodePools) == 1 {
		msg.Info("Cannot delete the last worker node pool in cluster")
		os.Exit(0)
	}

	nodePoolID := input.GetSelect("Node Pool ID:", "", nodePools, survey.Required)

	knpReq := &cloud.KubernetesNodePoolRequest{
		AccountID:           kc.AccountID,
		KubernetesClusterID: kc.KubernetesClusterID,
		ServiceID:           kc.ServiceID,
		ProviderID:          kc.ProviderID,
		NodePool: &cloud.KubernetesNodePool{
			NodePoolID: nodePoolID,
		},
	}

	output.ConfirmDeletion()

	s := output.Spinner()

	kc, err := nxc.DeleteCloudKubernetesNodePool(context.TODO(), knpReq)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete kubernetes node pool")
	}

	s.Stop()

	Output().Show(kc)
}
