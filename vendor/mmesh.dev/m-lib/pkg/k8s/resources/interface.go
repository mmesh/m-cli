package resources

import (
	"mmesh.dev/m-lib/pkg/k8s/resources/clusterrolebinding"
	"mmesh.dev/m-lib/pkg/k8s/resources/daemonset"
	"mmesh.dev/m-lib/pkg/k8s/resources/deployment"
	"mmesh.dev/m-lib/pkg/k8s/resources/namespace"
	"mmesh.dev/m-lib/pkg/k8s/resources/pod"
	"mmesh.dev/m-lib/pkg/k8s/resources/secret"
	"mmesh.dev/m-lib/pkg/k8s/resources/service"
	"mmesh.dev/m-lib/pkg/k8s/resources/serviceaccount"
	"mmesh.dev/m-lib/pkg/k8s/resources/statefulset"
)

type Interface interface {
	ServiceAccount() serviceaccount.Interface
	ClusterRoleBinding() clusterrolebinding.Interface
	Namespace() namespace.Interface
	Secret() secret.Interface
	StatefulSet() statefulset.Interface
	Deployment() deployment.Interface
	DaemonSet() daemonset.Interface
	Service() service.Interface
	Pod() pod.Interface
}

type api struct {
	kubeConfig []byte
}

func API(kubeConfig []byte) Interface {
	return &api{kubeConfig: kubeConfig}
}

func (a *api) ServiceAccount() serviceaccount.Interface {
	return &serviceaccount.API{KubeConfig: a.kubeConfig}
}

func (a *api) ClusterRoleBinding() clusterrolebinding.Interface {
	return &clusterrolebinding.API{KubeConfig: a.kubeConfig}
}

func (a *api) Namespace() namespace.Interface {
	return &namespace.API{KubeConfig: a.kubeConfig}
}

func (a *api) Secret() secret.Interface {
	return &secret.API{KubeConfig: a.kubeConfig}
}

func (a *api) StatefulSet() statefulset.Interface {
	return &statefulset.API{KubeConfig: a.kubeConfig}
}

func (a *api) Deployment() deployment.Interface {
	return &deployment.API{KubeConfig: a.kubeConfig}
}

func (a *api) DaemonSet() daemonset.Interface {
	return &daemonset.API{KubeConfig: a.kubeConfig}
}

func (a *api) Service() service.Interface {
	return &service.API{KubeConfig: a.kubeConfig}
}

func (a *api) Pod() pod.Interface {
	return &pod.API{}
}
