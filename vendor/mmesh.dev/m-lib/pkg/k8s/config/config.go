package config

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"mmesh.dev/m-lib/pkg/errors"
)

func NewClient(kubeConfig []byte) (*kubernetes.Clientset, error) {
	if kubeConfig != nil {
		return loadKubeConfig(kubeConfig)
	}

	return inClusterConfig()
}

func loadKubeConfig(kubeConfig []byte) (*kubernetes.Clientset, error) {
	apiConfig, err := clientcmd.Load(kubeConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function clientcmd.Load()", errors.Trace())
	}

	if err := clientcmd.Validate(*apiConfig); err != nil {
		return nil, errors.Wrapf(err, "[%v] function clientcmd.Validate()", errors.Trace())
	}

	configOverrides := &clientcmd.ConfigOverrides{}

	cConfig := clientcmd.NewDefaultClientConfig(*apiConfig, configOverrides)

	config, err := cConfig.ClientConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function cConfig.ClientConfig()", errors.Trace())
	}

	return kubernetes.NewForConfig(config)
}

func inClusterConfig() (*kubernetes.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function rest.InClusterConfig()", errors.Trace())
	}

	return kubernetes.NewForConfig(config)
}
