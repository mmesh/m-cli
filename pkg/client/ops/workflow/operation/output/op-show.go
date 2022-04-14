package output

import (
	"fmt"
	"strings"
	"time"

	"mmesh.dev/m-api-go/grpc/network/mmsp/command"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(o *workflow.Operation) {
	output.SectionHeader("Ops: Action Details")
	output.TitleT1("Action Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(o.AccountID))
	t.AddRow(colors.Black("Project ID"), colors.DarkWhite(o.ProjectID))
	t.AddRow(colors.Black("Workflow ID"), colors.DarkWhite(o.WorkflowID))
	t.AddRow(colors.Black("Job"), colors.DarkWhite(o.JobName))
	t.AddRow(colors.Black("Task"), colors.DarkWhite(o.TaskName))
	// t.AddRow("Operation ID:", colors.DarkWhite(o.OperationID), colors.DarkWhite(" "))
	// ref := strings.Split(o.OperationID, "@")[0]
	// actionName := strings.Split(ref, ":")[0]
	t.AddRow(colors.Black("Action Name"), colors.DarkWhite(o.ActionName))
	// t.AddRow("Target Node:", colors.DarkWhite(o.NodeID))
	// t.AddRow()
	// t.AddRow(colors.DarkWhite("Target"))
	t.AddRow(colors.Black("Tenant"), colors.DarkWhite(strings.Split(o.NodeID, ":")[1]))
	t.AddRow(colors.Black("Network"), colors.DarkWhite(strings.Split(o.NodeID, ":")[2]))
	t.AddRow(colors.Black("Subnet"), colors.DarkWhite(strings.Split(o.NodeID, ":")[3]))
	t.AddRow(colors.Black("Node"), colors.DarkWhite(strings.Split(o.NodeID, ":")[4]))

	// t.AddRow("Description:", colors.DarkWhite(o.Description))

	if o.Result != nil {
		switch o.Result.Status {
		case command.CommandResultStatus_EXECUTED:
			t.AddRow(colors.Black("Result"), output.StrOk("EXECUTED"))
		case command.CommandResultStatus_FAILED:
			t.AddRow(colors.Black("Result"), output.StrError("FAILED"))
		}
	}
	if o.Status != nil {
		tm := time.Unix(o.Status.Timestamp, 0)
		t.AddRow(colors.Black("Timestamp"), colors.DarkWhite(tm.String()))
	}

	t.Render()
	fmt.Println()

	if len(o.Stdout) > 0 || len(o.Stderr) > 0 || len(o.StdoutStderr) > 0 {
		output.SubTitleT2("Activity Log")

		if len(o.StdoutStderr) > 0 {
			fmt.Println(colors.Black("-----BEGIN OUTPUT-----"))
			fmt.Printf("%s", o.StdoutStderr)
			fmt.Println(colors.Black("-----END OUTPUT-----"))
			fmt.Println()
		} else {
			if len(o.Stdout) > 0 {
				fmt.Println(colors.Black("-----BEGIN STDOUT-----"))
				fmt.Printf("%s", o.Stdout)
				fmt.Println(colors.Black("-----END STDOUT-----"))
				fmt.Println()
			}

			if len(o.Stderr) > 0 {
				fmt.Println(colors.Black("-----BEGIN STDERR-----"))
				fmt.Printf("%s", o.Stderr)
				fmt.Println(colors.Black("-----END STDERR-----"))
				fmt.Println()
			}
		}
	}
}
