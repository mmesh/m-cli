package output

import (
	"fmt"

	"mmesh.dev/m-cli/pkg/client/k8s/resource"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(k8sResources map[string]*resource.KubernetesResource) {
	output.SectionHeader("Kubernetes: Resources")
	output.TitleT1("Resource List")

	t := table.New()

	t.SetColWidth(78)

	t.Header(colors.Black("KUBERNETES RESOURCE"))

	t.SetRowLine("-")

	for _, r := range k8sResources {
		var status string

		if r.Connected {
			// status = output.StrEnabled("■")
			status = colors.Green("█")
		} else {
			// status = output.StrDisabled("■")
			status = colors.DarkRed("█")
		}

		name := colors.DarkWhite(output.Fit(r.Name, 48))

		netStatus := getNetStatus(r)
		c1 := fmt.Sprintf("%s [%s] %s", status, r.Namespace, name)
		c1 = fmt.Sprintf("%s\n%s %s", c1, status, netStatus)

		t.AddRow(c1)
	}

	t.Render()
	fmt.Println()
}

func getNetStatus(r *resource.KubernetesResource) string {
	tenantID := r.TenantID
	if len(tenantID) == 0 {
		tenantID = "-"
	}
	netID := r.NetID
	if len(netID) == 0 {
		netID = "-"
	}
	vrfID := r.VRFID
	if len(vrfID) == 0 {
		vrfID = "-"
	}

	t := colors.Black("Tenant:")
	n := colors.Black("Network:")
	s := colors.Black("Subnet:")

	return fmt.Sprintf(
		"%s %s | %s %s | %s %s",
		t, colors.DarkGreen(tenantID), n, colors.DarkGreen(netID), s, colors.DarkGreen(vrfID))
}
