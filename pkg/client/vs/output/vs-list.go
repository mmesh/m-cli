package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(vss map[string]*topology.VS) {
	output.SectionHeader("Virtual Servers")
	output.TitleT1("VS List")

	t := table.New()
	t.Header(colors.Black("VS ID"), colors.Black("NAME"), colors.Black("DESCRIPTION"))

	for _, vs := range vss {
		t.AddRow(vs.VSID, colors.DarkWhite(vs.Name), output.Fit(vs.Description, 32))
		// t.AddRow(
		// 	vs.VSID,
		// 	colors.DarkWhite(vs.Name),
		// 	colors.DarkWhite(vs.Proto.String()),
		// 	colors.DarkWhite(fmt.Sprintf("%d", vs.VSPort)),
		// )
	}

	t.Render()
	fmt.Println()
}
