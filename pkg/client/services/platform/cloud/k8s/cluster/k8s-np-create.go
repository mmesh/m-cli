package cluster

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/client/provider"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
)

func (api *API) CreateNodePool() {
	kc := GetKubernetesCluster(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	vars.ProviderID = kc.ProviderID
	p := provider.GetProvider(services.ProviderType_CLOUD)
	nodeInstanceType := provider.GetKubernetesNodeInstanceType(p)

	it, err := getNodePoolInstanceType(p, nodeInstanceType.InstanceTypeID)
	if err != nil {
		status.Error(err, "Unable to find instanceType for cluster nodePool")
	}

	nodePoolID := input.GetInput("Node Pool ID:", "", "", input.ValidID)

	nNodes := input.GetInput("Nodes:", "", "3", input.ValidUint)

	numNodes, err := strconv.Atoi(nNodes)
	if err != nil {
		status.Error(errors.New("invalid number"), "invalid input")
	}

	knpReq := &cloud.KubernetesNodePoolRequest{
		AccountID:           kc.AccountID,
		KubernetesClusterID: kc.KubernetesClusterID,
		ServiceID:           kc.ServiceID,
		ProviderID:          kc.ProviderID,
		NodePool: &cloud.KubernetesNodePool{
			NodePoolID:       nodePoolID,
			NodeInstanceType: nodeInstanceType,
			NumNodes:         int32(numNodes),
			AutoScale:        false,
			MinNodes:         0,
			MaxNodes:         0,
		},
		InstanceType: it,
	}

	s := output.Spinner()

	kc, err = nxc.CreateCloudKubernetesNodePool(context.TODO(), knpReq)
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
