package output

import "mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"

type Interface interface {
	List(k8sClusters map[string]*cloud.KubernetesCluster)
	Show(kc *cloud.KubernetesCluster)
	KubeConfig(kc *cloud.KubernetesCluster)
}
type API struct{}
