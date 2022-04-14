package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(kc *cloud.KubernetesCluster) {
	output.SectionHeader("Cloud: Kubernetes Cluster Details")
	output.TitleT1("Kubernetes Cluster Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(kc.AccountID))
	// t.AddRow(colors.Black("Provider Cluster ID"), colors.DarkWhite(kc.ProviderKubernetesClusterID))
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

	output.SubTitleT2("Node Pools")

	t = table.New()

	i := 0
	for _, np := range kc.Spec.NodePools {
		i++
		knpID := fmt.Sprintf("%s%s%s Kubernetes Node Pool ID", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%d", i)), colors.DarkMagenta("]"))
		t.AddRow(colors.Black(knpID), colors.DarkWhite(np.ProviderNodePoolID))
		t.AddRow(colors.Black("    Node Pool Name"), colors.DarkWhite(np.Name))
		t.AddRow(colors.Black("    Instance Type"), colors.DarkWhite(np.NodeInstanceType.InstanceTypeID))
		t.AddRow(colors.Black("    Nodes"), colors.DarkWhite(fmt.Sprintf("%d", np.NumNodes)))
		t.AddRow("", "")

		if len(np.Nodes) > 0 {
			t.AddRow(colors.White("    Nodes"), "")
			t.AddRow("    =====", "")
			j := 0
			for _, node := range np.Nodes {
				j++
				nName := fmt.Sprintf("    %s%s%s Kubernetes Node ID", colors.DarkCyan("["), colors.Cyan(fmt.Sprintf("%d", j)), colors.DarkCyan("]"))
				t.AddRow(colors.Black(nName), colors.DarkWhite(node.KubernetesNodeID))
				t.AddRow(colors.Black("        Node Name"), colors.DarkWhite(node.Name))
				t.AddRow(colors.Black("        Instance ID"), colors.DarkWhite(node.InstanceID))

				if node.Managed {
					t.AddRow(colors.Black("        Features"), output.StrNormal("managed"))
				} else {
					t.AddRow(colors.Black("        Features"), output.StrDisabled("unmanaged"))
				}

				t.AddRow(colors.Black("        Status"), k8sStatus(node.Status))

				tm := time.Unix(node.CreationDate, 0).String()
				t.AddRow(colors.Black("        Creation Date"), colors.DarkWhite(tm))
				tm = time.Unix(node.LastModified, 0).String()
				t.AddRow(colors.Black("        Last Modified"), colors.DarkWhite(tm))

				t.AddRow("", "")
			}
		}
	}

	t.Render()
	fmt.Println()
}

func k8sStatus(s string) string {
	switch s {
	case "provisioning":
		return output.StrNormal(s)
	case "running":
		return output.StrOk(s)
	case "degraded":
		return output.StrWarning(s)
	case "error":
		return output.StrError(s)
	case "deleted":
		return output.StrInactive(s)
	case "upgrading":
		return output.StrNormal(s)
	case "invalid":
		return output.StrError(s)
	}

	return output.StrNormal(s)
}
