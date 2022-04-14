package cluster

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) AddNode() {
	kc := GetKubernetesCluster(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	var nodePoolID string
	var npOpts []string
	nodePools := make(map[string]*cloud.KubernetesNodePool)

	for _, np := range kc.Spec.NodePools {
		npOpts = append(npOpts, np.NodePoolID)
		nodePools[np.NodePoolID] = np
	}

	sort.Strings(npOpts)

	if len(npOpts) == 0 {
		msg.Alert("No node pool found")
		os.Exit(1)
	}

	nodePoolID = input.GetSelect("Node Pool ID:", "", npOpts, survey.Required)

	np := nodePools[nodePoolID]

	nNodes := input.GetInput("New Nodes:", "", "1", input.ValidUint)

	numNodes, err := strconv.Atoi(nNodes)
	if err != nil {
		status.Error(errors.New("invalid number"), "invalid input")
	}

	np.NumNodes += int32(numNodes)

	knpReq := &cloud.KubernetesNodePoolRequest{
		AccountID:           kc.AccountID,
		KubernetesClusterID: kc.KubernetesClusterID,
		ServiceID:           kc.ServiceID,
		ProviderID:          kc.ProviderID,
		NodePool:            np,
	}

	s := output.Spinner()

	kc, err = nxc.UpdateCloudKubernetesNodePool(context.TODO(), knpReq)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create cloud instance")
	}

	s.Stop()

	Output().Show(kc)

	fmt.Println(`Your kubernetes cluster is getting ready, but the process might take
5-10 minutes.
You can check the status with 'mmeshctl cloud kubernetes show' :-)`)
}
