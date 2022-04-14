package k8s

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/k8s/resource"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/k8s"
	"mmesh.dev/m-lib/pkg/k8s/config"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) ConnectService() {
	v := vrf.GetVRF(false)
	if v == nil {
		msg.Alert("No subnet found.")
		msg.Alert("Please, configure at least one.")
		os.Exit(1)
	}

	if api.kubeConfig == nil {
		kubeConfig, err := getKubeConfig()
		if err != nil {
			status.Error(err, "Unable to find KUBECONFIG")
		}
		api.kubeConfig = kubeConfig
	}

	mgw := api.validKubernetesGateway(v, api.getKubernetesServices())

	s := output.Spinner()

	resources, allIDs := api.getK8sResourceList(api.getKubernetesServicesAnnotations(), false)

	s.Stop()

	if len(allIDs) == 0 {
		msg.Info("All services already connected")
		os.Exit(1)
	}

	var selectedIDs []string

	selectMsg := fmt.Sprintf("Connect to %s", v.VRFID)
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

		annt := newAnnotations(r, v)
		if err := k8s.API(api.kubeConfig).Resources().Service().InsertAnnotations(r.Namespace, r.Name, annt); err != nil {
			status.Error(err, "Unable to set kubernetes service annotation")
		}
	}

	s.Stop()

	fmt.Println()

	api.Services()

	if !mgw {
		mgwNode := colors.DarkWhite(fmt.Sprintf("mgw-%s-%s-...", v.VRFID, v.NetID))
		fmt.Printf(`A new mmesh ingress gateway was created in the namespace %s.

If want to connect more kubernetes services to other subnets,
it is highly recommended that you use that namespace for your
mmesh gateways to keep things tidy.

You will see the gateway as %s in
the node list.

It will route all the services you want to connect to %s
in your kubernetes cluster.

Enjoy :-)
`, colors.DarkWhite(config.NamespaceDefault), mgwNode, colors.DarkWhite(v.VRFID))

		fmt.Println()
	}
}

func newAnnotations(r *resource.KubernetesResource, v *network.VRF) map[string]string {
	dnsName, ok := r.Annotations["mmesh.io/dnsName"]
	if !ok {
		dnsName = r.Name
	}
	ipv4, ok := r.Annotations["mmesh.io/ipv4"]
	if !ok {
		ipv4 = "auto"
	}

	return map[string]string{
		"mmesh.io/account": v.AccountID,
		"mmesh.io/tenant":  v.TenantID,
		"mmesh.io/network": v.NetID,
		"mmesh.io/subnet":  v.VRFID,
		"mmesh.io/dnsName": dnsName,
		"mmesh.io/ipv4":    ipv4,
	}
}
