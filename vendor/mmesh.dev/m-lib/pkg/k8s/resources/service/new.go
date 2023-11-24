package service

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/k8s/config"
)

func (a *API) New(i interface{}, appLabel config.AppLabel) *corev1.Service {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ni.K8SOpts.Name,
			Namespace: ni.K8SOpts.Ns,
			Labels: map[string]string{
				// "mmesh-federation": ni.FederationID,
				string(appLabel):   ni.K8SOpts.Name,
				"mmesh-app":        string(appLabel),
				"mmesh-objectID":   ni.K8SOpts.Name,
				"mmesh-type":       config.GetInstanceLabelType(ni.Type),
				"mmesh-account":    ni.Node.AccountID,
				"mmesh-tenant":     ni.Node.TenantID,
				"mmesh-network":    ni.Node.Cfg.NetID,
				"mmesh-subnet":     ni.Node.Cfg.SubnetID,
				"version":          ni.K8SOpts.Version,
			},
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				// "mmesh-federation": ni.FederationID,
				string(appLabel):   ni.K8SOpts.Name,
				"mmesh-app":        string(appLabel),
				"mmesh-objectID":   ni.K8SOpts.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:     "mmesh-p2p-tcp",
					Protocol: corev1.ProtocolTCP,
					Port:     ni.ExternalPort,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: ni.Node.Agent.Port,
					},
					NodePort: ni.ExternalPort,
				},
				{
					Name:     "mmesh-p2p-quic",
					Protocol: corev1.ProtocolUDP,
					Port:     ni.ExternalPort,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: ni.Node.Agent.Port,
					},
					NodePort: ni.ExternalPort,
				},
			},
			SessionAffinity:       corev1.ServiceAffinityClientIP,
			ExternalTrafficPolicy: corev1.ServiceExternalTrafficPolicyTypeLocal,
			// LoadBalancerIP:        ni.ExternalIPv4,
			// LoadBalancerSourceRanges: []string{
			// 	"0.0.0.0/0",
			// },
		},
	}
}

/*
func (a *API) NewDynamic() *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      r.NodeInstance.K8SOpts.Name,
				"namespace": r.NodeInstance.K8SOpts.Ns,
				"labels": map[string]interface{}{
					"mmesh-federation": r.NodeInstance.FederationID,
					"mmesh-node":       r.NodeInstance.K8SOpts.Name,
					"mmesh-app":        "mmesh-node",
					"mmesh-objectID":   r.NodeInstance.K8SOpts.Name,
					"version":          r.NodeInstance.Version,
				},
			},
			"spec": map[string]interface{}{
				"type": "NodePort",
				"selector": map[string]interface{}{
					"mmesh-federation": r.NodeInstance.FederationID,
					"mmesh-node":       r.NodeInstance.K8SOpts.Name,
					"mmesh-app":        "mmesh-node",
					"mmesh-objectID":   r.NodeInstance.K8SOpts.Name,
				},
				"ports": []map[string]interface{}{
					{
						"name":       "mmesh-p2p",
						"protocol":   "TCP",
						"port":       r.NodeInstance.ExternalPort,
						"targetPort": r.NodeInstance.Agent.Node.Agent.Port,
						"nodePort":   r.NodeInstance.ExternalPort,
					},
				},
				"sessionAffinity":       "ClientIP",
				"externalTrafficPolicy": "Local",
				// "loadBalancerIP":        r.NodeInstance.ExternalIPv4,
				// "loadBalancerSourceRanges": []string{
				// 	"0.0.0.0/0",
				// },
			},
		},
	}
}
*/
