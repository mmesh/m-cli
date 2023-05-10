package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(projects map[string]*ops.Project) {
	output.SectionHeader("Ops: Projects")
	output.TitleT1("Project List")

	t := table.New()
	t.Header(colors.Black("PROJECT NAME"), colors.Black("DESCRIPTION"))

	for _, p := range projects {
		t.AddRow(colors.DarkWhite(output.Fit(p.Name, 24)), output.Fit(p.Description, 40))
	}

	t.Render()
	fmt.Println()
}
