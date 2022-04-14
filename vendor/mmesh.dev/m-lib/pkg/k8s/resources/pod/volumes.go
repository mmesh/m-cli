package pod

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/k8s/config"
)

func (a *API) NewVolumes(i interface{}, appLabel config.AppLabel) []corev1.Volume {
	var ni *network.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*network.NodeInstance)
	default:
		return nil
	}

	configVolumeName := fmt.Sprintf("%s-config", ni.K8SOpts.Name)

	return []corev1.Volume{
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
	}
}
