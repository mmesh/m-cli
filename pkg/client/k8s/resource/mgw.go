package resource

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
)

func (r *KubernetesResource) GetGatewayNodeInstance(s *topology.Subnet, externalPort int32) (*topology.NodeInstance, error) {
	if len(r.Name) == 0 || len(r.Namespace) == 0 {
		return nil, fmt.Errorf("missing kubernetes resource metadata")
	}

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	req := &topology.NewK8SGatewayRequest{
		NodeRequest: &topology.NewNodeRequest{
			AccountID:   s.AccountID,
			TenantID:    s.TenantID,
			NetID:       s.NetID,
			SubnetID:    s.SubnetID,
			NodeName:    r.Name,
			Description: fmt.Sprintf("[k8s-gw] %s", r.Name),
			Port:        externalPort,
			Type:        topology.NodeType_K8S_GATEWAY,
		},
		ExternalPort: externalPort,
		K8SOptsReq: &topology.KubernetesOptsRequest{
			Namespace:  r.Namespace,
			Name:       r.Name,
			ReplicaSet: false,
		},
	}

	return nxc.CreateKubernetesGateway(context.TODO(), req)
}
