package output

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/utils"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(operations []*workflow.Operation) {
	output.SectionHeader("Ops: Actions")
	output.TitleT1("Action List")

	t := table.New()
	t.Header(colors.Black("ACTION"), colors.Black("TIMESTAMP"), colors.Black("TARGET NODE"))

	for _, o := range operations {
		ref := strings.Split(o.OperationID, "@")[0]
		name := strings.Split(ref, ":")[0]
		tm, _ := utils.ParseTimestamp(strings.Split(ref, ":")[1])
		// mmid := strings.Split(o.OperationID, "@")[1]
		// nodeID := output.Fit(strings.Split(mmid, ":")[4], 32)
		mmID := mmid.MMNodeID(o.NodeID)
		nodeID := output.Fit(mmID.NodeID(), 32)
		// oID := fmt.Sprintf("%s", colors.DarkWhite(o.OperationID))
		t.AddRow(colors.DarkWhite(name), output.Datetime(tm.Unix()), nodeID)
	}

	t.Render()
	fmt.Println()
}
