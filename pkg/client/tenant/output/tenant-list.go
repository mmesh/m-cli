package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(tenants map[string]*tenant.Tenant) {
	output.SectionHeader("Tenants")
	output.TitleT1("Tenant List")

	t := table.New()
	t.Header(colors.Black("TENANT ID"), colors.Black("DESCRIPTION"))

	for _, tenant := range tenants {
		t.AddRow(colors.DarkWhite(tenant.TenantID), output.Fit(tenant.Description, 48))
	}

	t.Render()
	fmt.Println()
}
