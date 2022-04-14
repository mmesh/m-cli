package resource

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/grpc"
)

func (r *KubernetesResource) GetPodNodeInstance(v *network.VRF) (*network.NodeInstance, error) {
	if len(r.Name) == 0 || len(r.Namespace) == 0 {
		return nil, fmt.Errorf("missing kubernetes resource metadata")
	}

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	niReq := &network.NodeInstanceRequest{
		AccountID: v.AccountID,
		TenantID:  v.TenantID,
		NetID:     v.NetID,
		VRFID:     v.VRFID,
		NodeID:    r.Name,
		// ExternalPort: ,
		K8SOptsReq: &network.KubernetesOptsRequest{
			Namespace:  r.Namespace,
			Name:       fmt.Sprintf("%s-mmesh-node", r.Name),
			ReplicaSet: true,
		},
	}

	return nxc.CreateGenericNode(context.TODO(), niReq)
}
