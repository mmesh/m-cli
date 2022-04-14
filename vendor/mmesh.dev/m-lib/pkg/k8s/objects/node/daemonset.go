package node

import (
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/k8s/config"
	"mmesh.dev/m-lib/pkg/k8s/resources"
)

func (a *API) ConnectDaemonSet(ns, name string, ni *network.NodeInstance) error {
	secret := resources.API(a.KubeConfig).Secret().New(ni, config.AppLabelNode)
	if _, err := resources.API(a.KubeConfig).Secret().Create(secret); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Secret().Create()", errors.Trace())
	}

	c := resources.API(a.KubeConfig).Pod().NewContainer(ni, config.AppLabelNode)
	vols := resources.API(a.KubeConfig).Pod().NewVolumes(ni, config.AppLabelNode)
	labels := resources.API(a.KubeConfig).Pod().NewLabels(ni, config.AppLabelNode)
	if err := resources.API(a.KubeConfig).DaemonSet().AddContainer(ns, name, *c, vols, labels); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().DaemonSet().AddContainer()", errors.Trace())
	}

	return nil
}

func (a *API) DisconnectDaemonSet(ns, name string, ni *network.NodeInstance) error {
	c := resources.API(a.KubeConfig).Pod().NewContainer(ni, config.AppLabelNode)
	vols := resources.API(a.KubeConfig).Pod().NewVolumes(ni, config.AppLabelNode)
	labels := resources.API(a.KubeConfig).Pod().NewLabels(ni, config.AppLabelNode)
	if err := resources.API(a.KubeConfig).DaemonSet().RemoveContainer(ns, name, *c, vols, labels); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().DaemonSet().RemoveContainer()", errors.Trace())
	}

	secretName := ni.K8SOpts.Name

	if err := resources.API(a.KubeConfig).Secret().Delete(ns, secretName); err != nil {
		return errors.Wrapf(err, "[%v] function resources.API().Secret().Delete()", errors.Trace())
	}

	return nil
}
