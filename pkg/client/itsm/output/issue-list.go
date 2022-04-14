package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(issues map[string]*itsm.Issue) {
	output.SectionHeader("ITSM: Issues")
	output.TitleT1("Issue List")

	if len(issues) == 0 {
		fmt.Println(colors.Black("No data"))
		fmt.Println()
		return
	}

	t := table.New()

	// t.SetColWidth(24)

	t.Header(colors.Black("UPDATED"), colors.Black("SUMMARY"), colors.Black("STATUS"))

	// t.SetAutoWrapText(true)
	// t.SetReflowDuringAutoWrap(true)

	// t.SetRowLine("─")
	t.SetRowLine("-")

	for _, i := range issues {
		tm := output.Datetime(i.LastModified)
		status := output.StrFixedField(issueStatus(i.Status))
		summary := output.Fit(i.Summary, 36)
		t.AddRow(colors.DarkWhite(tm), summary, status)
	}

	t.Render()
	fmt.Println()
}
