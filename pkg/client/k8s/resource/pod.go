package resource

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
)

func (r *KubernetesResource) GetPodNodeInstance(s *topology.Subnet) (*topology.NodeInstance, error) {
	if len(r.Name) == 0 || len(r.Namespace) == 0 {
		return nil, fmt.Errorf("missing kubernetes resource metadata")
	}

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	req := &topology.NewK8SPodRequest{
		NodeRequest: &topology.NewNodeRequest{
			AccountID:   s.AccountID,
			TenantID:    s.TenantID,
			NetID:       s.NetID,
			SubnetID:    s.SubnetID,
			NodeName:    r.Name,
			Description: fmt.Sprintf("[k8s-pod] %s", r.Name),
			Type:        topology.NodeType_K8S_POD,
		},
		K8SOptsReq: &topology.KubernetesOptsRequest{
			Namespace:  r.Namespace,
			Name:       fmt.Sprintf("%s-mmesh-node", r.Name),
			ReplicaSet: true,
		},
	}

	return nxc.CreateKubernetesPod(context.TODO(), req)
}
