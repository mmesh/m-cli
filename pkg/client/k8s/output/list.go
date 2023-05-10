package output

import (
	"fmt"

	tenant_pb "mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/client/k8s/resource"
	"mmesh.dev/m-cli/pkg/client/tenant"
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

	s := output.Spinner()

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

	s.Stop()

	t.Render()
	fmt.Println()
}

func getNetStatus(r *resource.KubernetesResource) string {
	tenantName := "-"
	if len(r.TenantID) > 0 {
		t := getTenant(r.TenantID)
		if t != nil {
			tenantName = t.Name
		}
	}

	netID := r.NetID
	if len(netID) == 0 {
		netID = "-"
	}
	subnetID := r.SubnetID
	if len(subnetID) == 0 {
		subnetID = "-"
	}

	t := colors.Black("Tenant:")
	n := colors.Black("Network:")
	s := colors.Black("Subnet:")

	return fmt.Sprintf(
		"%s %s | %s %s | %s %s",
		t, colors.DarkGreen(tenantName), n, colors.DarkGreen(netID), s, colors.DarkGreen(subnetID))
}

var tenantsMap = make(map[string]*tenant_pb.Tenant, 0)

func getTenant(tenantID string) *tenant_pb.Tenant {
	t, ok := tenantsMap[tenantID]
	if ok {
		return t
	}

	t = tenant.FetchTenant(tenantID)
	if t != nil {
		tenantsMap[tenantID] = t
		return t
	}

	return nil
}
