package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/ops/project"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(projects map[string]*project.Project) {
	output.SectionHeader("Ops: Projects")
	output.TitleT1("Project List")

	t := table.New()
	t.Header(colors.Black("PROJECT ID"), colors.Black("DESCRIPTION"))

	for _, p := range projects {
		t.AddRow(colors.DarkWhite(p.ProjectID), output.Fit(p.Description, 48))
	}

	t.Render()
	fmt.Println()
}
