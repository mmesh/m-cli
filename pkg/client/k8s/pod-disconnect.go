package k8s

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/k8s/resource"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/k8s"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) DisconnectPod() {
	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	s := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesPods(), true)

	s.Stop()

	if len(allIDs) == 0 {
		msg.Info("No pods connected found")
		os.Exit(1)
	}

	var selectedIDs []string

	selectMsg := "Disconnect from mmesh"
	if err := survey.AskOne(
		&survey.MultiSelect{
			Message:  selectMsg,
			Options:  allIDs,
			PageSize: 10,
		},
		&selectedIDs,
		survey.WithIcons(input.SurveySetIcons),
	); err != nil {
		status.Error(err, "Unable to get response")
	}

	s = output.Spinner()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		if !r.Connected {
			continue
		}

		ni, err := r.GetPodNodeInstance(&network.VRF{
			AccountID: r.AccountID,
			TenantID:  r.TenantID,
			NetID:     r.NetID,
			VRFID:     r.VRFID,
		})
		if err != nil {
			status.Error(err, "Unable to get node instance")
		}

		switch r.KubernetesResourceType {
		case resource.KubernetesResourceTypeStatefulSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectStatefulSet(r.Namespace, r.Name, ni); err != nil {
				status.Error(err, "Unable to disconnect kubernetes statefulSet")
			}
		case resource.KubernetesResourceTypeDeployment:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectDeployment(r.Namespace, r.Name, ni); err != nil {
				status.Error(err, "Unable to disconnect kubernetes deployment")
			}
		case resource.KubernetesResourceTypeDaemonSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectDaemonSet(r.Namespace, r.Name, ni); err != nil {
				status.Error(err, "Unable to disconnect kubernetes daemonSet")
			}
		}
	}

	s.Stop()

	fmt.Println()

	api.Pods()
}
