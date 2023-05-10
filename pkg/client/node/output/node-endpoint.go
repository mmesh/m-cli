package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) ShowEndpoint(e *topology.Endpoint) {
	output.SectionHeader("Endpoint Details")
	output.TitleT1("Network Endpoint")

	t := table.New()

	t.AddRow(colors.Black("Endpoint"), colors.DarkWhite(e.EndpointID))
	t.AddRow(colors.Black("FQDN"), colors.DarkWhite(e.DNSName+".mmesh.local"))
	t.AddRow(colors.Black("IPv4"), colors.DarkWhite(e.IPv4))
	t.AddRow(colors.Black("IPv6"), colors.DarkWhite(e.IPv6))

	t.Render()
	fmt.Println()
}
