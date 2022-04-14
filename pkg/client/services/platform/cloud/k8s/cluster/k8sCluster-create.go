package cluster

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/client/provider"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
)

func (api *API) Create() {
	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	v := vrf.GetVRF(false)

	p := provider.GetProvider(services.ProviderType_CLOUD)
	kOpts := provider.GetKubernetesOptions(p)

	if kOpts == nil {
		status.Error(fmt.Errorf("Kubernetes not supported yet by %s integration", p.Name), "Unable to create resource")
	}

	it, err := getNodePoolInstanceType(p, kOpts.NodeInstanceTypes[0].InstanceTypeID)
	if err != nil {
		status.Error(err, "Unable to find instanceType for cluster default nodePool")
	}

	vars.CloudKubernetesClusterID = input.GetInput("Cluster ID:", "", vars.CloudKubernetesClusterID, input.ValidID)

	nNodes := input.GetInput("Nodes:", "", "3", input.ValidUint)

	numNodes, err := strconv.Atoi(nNodes)
	if err != nil {
		status.Error(errors.New("invalid number"), "invalid input")
	}

	kcReq := &cloud.KubernetesClusterRequest{
		AccountID:           v.AccountID,
		VRF:                 v,
		KubernetesClusterID: vars.CloudKubernetesClusterID,
		Region:              kOpts.Regions[0],
		Version:             kOpts.Versions[0],
		Spec: &cloud.KubernetesClusterSpec{
			NodePools: map[string]*cloud.KubernetesNodePool{
				"default": {
					NodePoolID:       "default",
					NodeInstanceType: kOpts.NodeInstanceTypes[0],
					NumNodes:         int32(numNodes),
					AutoScale:        false,
					MinNodes:         0,
					MaxNodes:         0,
				},
			},
			InstanceTypes: map[string]*compute.InstanceType{
				"default": it,
			},
			AutoUpgrade:  true,
			SurgeUpgrade: false,
		},
	}

	s := output.Spinner()

	kc, err := nxc.CreateCloudKubernetesCluster(context.TODO(), kcReq)
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

func getNodePoolInstanceType(p *services.Provider, itID string) (*compute.InstanceType, error) {
	for _, it := range p.Catalog.Cloud.InstanceTypes {
		if it.InstanceTypeID == itID {
			return it, nil
		}
	}

	return nil, fmt.Errorf("instanceType not found in catalog")
}
