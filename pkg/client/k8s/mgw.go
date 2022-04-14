package k8s

import (
	"fmt"
	"os"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/k8s/resource"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/k8s"
	"mmesh.dev/m-lib/pkg/k8s/config"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) validKubernetesGateway(v *network.VRF, gateways map[string]*resource.KubernetesResource) bool {
	for _, sgw := range gateways {
		if sgw.AccountID == v.AccountID &&
			sgw.TenantID == v.TenantID &&
			sgw.NetID == v.NetID &&
			sgw.VRFID == v.VRFID {
			// kubernetes gateway found for this subnet
			return true
		}
	}

	msg.Infof("No mmesh ingress gateway found to route the %s in your cluster.", colors.DarkWhite(v.VRFID))

	fmt.Println()

	mgwMsg := "Want to create one?"
	if input.GetConfirm(mgwMsg, true) {
		s := output.Spinner()
		api.createKubernetesGateway(v)
		s.Stop()

		return false
	} else {
		fmt.Println()
		fmt.Println("Exiting..")
		fmt.Println()

		os.Exit(0)
	}

	return false
}

func (api *API) createKubernetesGateway(v *network.VRF) {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	port, err := k8s.API(api.kubeConfig).Resources().Service().GetNodePort()
	if err != nil {
		status.Error(err, "Unable to get node port")
	}

	r := &resource.KubernetesResource{
		Namespace: config.NamespaceDefault,
		Name:      fmt.Sprintf("mgw-%s-%s-%s-%s-%d", v.VRFID, v.NetID, v.TenantID, v.AccountID, time.Now().Unix()),
	}

	ni, err := r.GetGatewayNodeInstance(v, int32(port))
	if err != nil {
		status.Error(err, "Unable to get node instance")
	}

	if err := k8s.API(api.kubeConfig).Objects().Node().CreateGateway(ni); err != nil {
		status.Error(err, "Unable to create kubernetes resources")
	}
}
