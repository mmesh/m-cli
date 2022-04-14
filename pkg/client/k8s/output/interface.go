package output

import "mmesh.dev/m-cli/pkg/client/k8s/resource"

type Interface interface {
	List(k8sResources map[string]*resource.KubernetesResource)
}
type API struct{}
