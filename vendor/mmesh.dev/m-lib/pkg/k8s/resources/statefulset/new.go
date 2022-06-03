package statefulset

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/k8s/config"
	"mmesh.dev/m-lib/pkg/mm"
)

func (a *API) New(i interface{}, appLabel config.AppLabel) *appsv1.StatefulSet {
	var ni *network.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*network.NodeInstance)
	default:
		return nil
	}

	configVolumeName := fmt.Sprintf("%s-config", string(appLabel))

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "StatefulSet",
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
		Spec: appsv1.StatefulSetSpec{
			ServiceName: ni.K8SOpts.Name,
			Replicas:    mm.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"mmesh-federation": ni.FederationID,
					string(appLabel):   ni.K8SOpts.Name,
					"mmesh-app":        string(appLabel),
					"mmesh-objectID":   ni.K8SOpts.Name,
				},
			},
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"mmesh-federation": ni.FederationID,
						string(appLabel):   ni.K8SOpts.Name,
						"mmesh-app":        string(appLabel),
						"mmesh-objectID":   ni.K8SOpts.Name,
						"version":          ni.K8SOpts.Version,
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: config.ServiceAccountView,
					Hostname:           ni.Node.NodeID,
					Containers: []corev1.Container{
						{
							Name:            ni.K8SOpts.Name,
							Image:           ni.K8SOpts.Image,
							ImagePullPolicy: corev1.PullAlways,
							SecurityContext: &corev1.SecurityContext{
								Privileged: mm.Bool(true), // only needed by sysctl to enable ipv6
								Capabilities: &corev1.Capabilities{
									Add: []corev1.Capability{
										"net_admin",
									},
								},
							},
							Args: []string{
								"start",
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "mmesh-p2p-tcp",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: ni.Node.Agent.Port,
								},
								{
									Name:          "mmesh-p2p-quic",
									Protocol:      corev1.ProtocolUDP,
									ContainerPort: ni.Node.Agent.Port,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "dev-net-tun",
									MountPath: "/dev/net/tun",
									ReadOnly:  true,
								},
								{
									Name:      configVolumeName,
									MountPath: "/etc/mmesh",
									ReadOnly:  true,
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									// cpu: 30m
									corev1.ResourceCPU: *resource.NewMilliQuantity(30, resource.DecimalSI),
									// memory: 100Mi
									corev1.ResourceMemory: *resource.NewQuantity(100*1024*1024, resource.BinarySI),
								},
								Requests: corev1.ResourceList{
									// cpu: 15m
									corev1.ResourceCPU: *resource.NewMilliQuantity(15, resource.DecimalSI),
									// memory: 50Mi
									corev1.ResourceMemory: *resource.NewQuantity(50*1024*1024, resource.BinarySI),
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "dev-net-tun",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/dev/net/tun",
								},
							},
						},
						{
							Name: configVolumeName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: ni.K8SOpts.Name,
								},
							},
						},
					},
				},
			},
		},
	}
}

/*
func (a *API) NewDynamic() *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
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
				"serviceName": r.NodeInstance.K8SOpts.Name,
				"replicas":    1,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"mmesh-federation": r.NodeInstance.FederationID,
						"mmesh-node":       r.NodeInstance.K8SOpts.Name,
						"mmesh-app":        "mmesh-node",
						"mmesh-objectID":   r.NodeInstance.K8SOpts.Name,
					},
				},
				"updateStrategy": map[string]interface{}{
					"type": "RollingUpdate",
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"mmesh-federation": r.NodeInstance.FederationID,
							"mmesh-node":       r.NodeInstance.K8SOpts.Name,
							"mmesh-app":        "mmesh-node",
							"mmesh-objectID":   r.NodeInstance.K8SOpts.Name,
							"version":          r.NodeInstance.Version,
						},
					},
					"spec": map[string]interface{}{
						// "serviceAccountName": "mmesh-view",
						"hostname": r.NodeInstance.Agent.Node.NodeID,
						"containers": []map[string]interface{}{
							{
								"name":            r.NodeInstance.K8SOpts.Name,
								"image":           r.NodeInstance.Image,
								"imagePullPolicy": "Always",
								"securityContext": map[string]interface{}{
									"privileged": true, // only needed by sysctl to enable ipv6
									"capabilities": map[string]interface{}{
										"add": []string{
											"net_admin",
										},
									},
								},
								"args": []string{
									"start",
								},
								"ports": []map[string]interface{}{
									{
										"name":          "mmesh-p2p",
										"protocol":      "TCP",
										"containerPort": r.NodeInstance.Agent.Node.Agent.Port,
									},
								},
								"volumeMounts": []map[string]interface{}{
									{
										"name":      "dev-net-tun",
										"mountPath": "/dev/net/tun",
										"readOnly":  true,
									},
									{
										"name":      "mmesh-node-config",
										"mountPath": "/etc/mmesh",
										"readOnly":  true,
									},
								},
							},
						},
						"volumes": []map[string]interface{}{
							{
								"name": "dev-net-tun",
								"hostPath": map[string]interface{}{
									"path": "/dev/net/tun",
								},
							},
							{
								"name": "mmesh-node-config",
								"secret": map[string]interface{}{
									"secretName": r.NodeInstance.K8SOpts.Name,
								},
							},
						},
					},
				},
			},
		},
	}
}
*/
