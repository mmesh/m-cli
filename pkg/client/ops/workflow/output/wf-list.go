package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/client/event"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(workflows map[string]*ops.Workflow) {
	output.SectionHeader("Ops: Workflows")
	output.TitleT1("Workflow List")

	t := table.New()
	t.Header(colors.Black("WORKFLOW NAME"), colors.Black(output.Fit("FAILURE RATE", 14)))

	for _, wf := range workflows {
		wfID := colors.DarkWhite(output.Fit(wf.Name, 48))
		if wf.EventMetrics != nil {
			t.AddRow(wfID, event.Output().FailureProbability(wf.EventMetrics))
		} else {
			t.AddRow(wfID, "n/a")
		}
	}

	t.Render()
	fmt.Println()
}
