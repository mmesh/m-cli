package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) KubeConfig(kc *cloud.KubernetesCluster) {
	output.SectionHeader("Cloud: Kubernetes Cluster Details")
	output.TitleT1("Kubernetes Cluster Information")

	t := table.New()

	t.AddRow(colors.Black("Cluster ID"), colors.DarkWhite(kc.KubernetesClusterID))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(kc.Description))
	t.AddRow(colors.Black("Region"), colors.DarkWhite(kc.Region))
	t.AddRow(colors.Black("Version"), colors.DarkWhite(kc.Version))

	var managed string
	if kc.Managed {
		managed = output.StrNormal("managed")
	} else {
		managed = output.StrDisabled("unmanaged")
	}
	var autoUpgrade string
	if kc.Spec.AutoUpgrade {
		autoUpgrade = output.StrNormal("autoUpgrade")
	}

	t.AddRow(colors.Black("Features"), managed+" "+autoUpgrade)

	t.AddRow(colors.Black("Cluster Subnet"), colors.DarkWhite(kc.ClusterSubnet))
	t.AddRow(colors.Black("Service Subnet"), colors.DarkWhite(kc.ServiceSubnet))

	t.AddRow(colors.Black("Status"), k8sStatus(kc.Status))

	tm := time.Unix(kc.CreationDate, 0).String()
	t.AddRow(colors.Black("Creation Date"), colors.DarkWhite(tm))
	tm = time.Unix(kc.LastModified, 0).String()
	t.AddRow(colors.Black("Last Modified"), colors.DarkWhite(tm))

	t.Render()
	fmt.Println()

	if len(kc.Kubeconfig) == 0 {
		fmt.Print(colors.DarkMagenta("[cluster is getting ready, please wait a couple of minutes]\n"))
	} else {
		fmt.Print(colors.Black("-----KUBECONFIG-----\n"))
		fmt.Printf("%s\n", string(kc.Kubeconfig))
		fmt.Print(colors.Black("-----KUBECONFIG-----\n"))
	}

	fmt.Println()
}
