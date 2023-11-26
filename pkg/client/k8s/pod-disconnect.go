package k8s

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/k8s/resource"
	"mmesh.dev/m-cli/pkg/grpc"
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

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	for _, rID := range selectedIDs {
		r, ok := resources[rID]
		if !ok {
			s.Stop()
			msg.Error("Unable to parse response")
			os.Exit(1)
		}

		if !r.Connected {
			continue
		}

		ni, err := r.GetPodNodeInstance(&topology.Subnet{
			AccountID: r.AccountID,
			TenantID:  r.TenantID,
			NetID:     r.NetID,
			SubnetID:  r.SubnetID,
		})
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to get node instance")
		}

		switch r.KubernetesResourceType {
		case resource.KubernetesResourceTypeStatefulSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectStatefulSet(r.Namespace, r.Name, ni); err != nil {
				s.Stop()
				status.Error(err, "Unable to disconnect kubernetes statefulSet")
			}
		case resource.KubernetesResourceTypeDeployment:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectDeployment(r.Namespace, r.Name, ni); err != nil {
				s.Stop()
				status.Error(err, "Unable to disconnect kubernetes deployment")
			}
		case resource.KubernetesResourceTypeDaemonSet:
			if err := k8s.API(api.kubeConfig).Objects().Node().DisconnectDaemonSet(r.Namespace, r.Name, ni); err != nil {
				s.Stop()
				status.Error(err, "Unable to disconnect kubernetes daemonSet")
			}
		}

		nr := &topology.NodeReq{
			AccountID: ni.Node.AccountID,
			TenantID:  ni.Node.TenantID,
			NodeID:    ni.Node.NodeID,
		}

		if _, err := nxc.DeleteNode(context.TODO(), nr); err != nil {
			s.Stop()
			status.Error(err, "Unable to delete node for k8s-pod")
		}
	}

	s.Stop()

	fmt.Println()

	api.Pods()
}
