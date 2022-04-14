package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(n *network.Network) {
	output.SectionHeader("Network Details")
	output.TitleT1("Network Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(n.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(n.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(n.NetID))
	t.AddRow(colors.Black("Network CIDR"), colors.DarkWhite(n.NetworkCIDR))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(n.Description))

	if n.RouteVRFs {
		t.AddRow(colors.Black("Routed Subnets"), output.StrEnabled("yes"))
	} else {
		t.AddRow(colors.Black("Routed Subnets"), output.StrDisabled("no"))
	}

	t.Render()
	fmt.Println()
}
