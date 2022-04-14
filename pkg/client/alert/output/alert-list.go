package output

import (
	"fmt"
	"sort"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) List(alerts map[string]*events.Alert) {
	output.SectionHeader("Alerts")
	output.TitleT1("Alert List")

	if len(alerts) == 0 {
		msg.Info("No records found")
		return
	}

	t := table.New()
	t.Header(colors.Black("LAST UPDATED / NODE ID"), colors.Black("STATUS / COMPONENT"))

	t.SetRowLine("-")

	tmSort := make([]string, 0)
	tmSortUniq := make(map[string]struct{})
	for _, a := range alerts {
		tm := output.Datetime(a.LastUpdated)
		if _, ok := tmSortUniq[tm]; !ok {
			tmSortUniq[tm] = struct{}{}
			tmSort = append(tmSort, tm)
		}
	}
	sort.Strings(tmSort)

	for _, tm := range tmSort {
		for _, a := range alerts {
			if tm == output.Datetime(a.LastUpdated) {
				nodeID := colors.DarkWhite(output.Fit(a.NodeID, 36))
				component := colors.DarkWhite(output.Fit(strings.ToLower(a.Component), 32))

				var status string
				switch strings.ToLower(a.Status) {
				case "triggered":
					status = colors.DarkRed("█")
				case "acknowledged":
					status = colors.DarkYellow("█")
				case "resolved":
					status = colors.Green("█")
				}

				alertStatus := api.AlertStatus(a.Status)

				c1 := fmt.Sprintf("%s %s\n%s %s", status, colors.Black(tm), status, nodeID)
				c2 := fmt.Sprintf("%s\n%s", alertStatus, component)

				t.AddRow(c1, c2)
			}
		}
	}

	t.Render()
	fmt.Println()
}
