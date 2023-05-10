package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(taskLogs []*ops.TaskLog) {
	output.SectionHeader("Ops: TaskLogs")
	output.TitleT1("TaskLog List")

	t := table.New()
	t.Header(colors.Black("TASK"), colors.Black("TIMESTAMP"), colors.Black("TARGET NODE"))

	for _, tl := range taskLogs {
		tm := tl.Status.Timestamp
		t.AddRow(colors.DarkWhite(tl.TaskName), output.DatetimeMilli(tm), tl.NodeName)
	}

	t.Render()
	fmt.Println()
}
