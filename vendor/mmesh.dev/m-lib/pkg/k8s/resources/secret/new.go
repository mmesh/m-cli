package secret

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/k8s/config"
)

func (a *API) New(i interface{}, appLabel config.AppLabel) *corev1.Secret {
	var ni *network.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*network.NodeInstance)
	default:
		return nil
	}

	ymlFile := fmt.Sprintf("%s.yml", string(appLabel))

	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ni.K8SOpts.Name,
			Namespace: ni.K8SOpts.Ns,
			Labels: map[string]string{
				"mmesh-federation": ni.FederationID,
				string(appLabel):   ni.K8SOpts.Name,
				"mmesh-app":        string(appLabel),
				"mmesh-objectID":   ni.K8SOpts.Name,
				"mmesh-type":       config.GetInstanceLabelType(ni.Type),
				"version":          ni.K8SOpts.Version,
			},
		},
		Type: corev1.SecretTypeOpaque,
		StringData: map[string]string{
			ymlFile: ni.Config.YAML,
		},
	}
}

/*
func (a *API) NewDynamic() *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      r.NodeInstance.K8SOpts.Name,
				"namespace": r.NodeInstance.K8SOpts.Ns,
				"labels": map[string]interface{}{
					"mmesh-federation": r.FederationID,
					"mmesh-node":       r.NodeInstance.K8SOpts.Name,
					"mmesh-app":        "mmesh-node",
					"mmesh-objectID":   r.NodeInstance.K8SOpts.Name,
					"version":          r.NodeInstance.K8SOpts.Version,
				},
			},
			"type": "Opaque",
			"stringData": map[string]interface{}{
				"mmesh-node.yml": r.Config.YAML,
			},
		},
	}
}
*/
