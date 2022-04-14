package output

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(a *events.Alert) {
	// fmt.Println() // extra line needed due to the long select option above
	output.SectionHeader("Alerts")
	output.TitleT1("Alert Information")

	t := table.New()
	t.BulkData([][]string{
		{colors.Black("Account ID"), colors.DarkWhite(a.AccountID)},
		{colors.Black("Alert ID"), colors.DarkWhite(output.Fit(a.AlertID, 62))},
		{colors.Black("Tenant ID"), colors.DarkWhite(a.TenantID)},
		{colors.Black("Network ID"), colors.DarkWhite(a.NetID)},
		{colors.Black("Subnet ID"), colors.DarkWhite(a.VRFID)},
		{colors.Black("Node ID"), colors.DarkWhite(a.NodeID)},
		{colors.Black("Severity"), alertSeverity(a.Severity)},
		{colors.Black("Component"), colors.DarkWhite(output.Fit(a.Component, 62))},
		{colors.Black("Group"), colors.DarkWhite(output.Fit(a.Group, 62))},
		{colors.Black("Class"), colors.DarkWhite(a.Class)},
		{colors.Black("Status"), api.AlertStatus(a.Status)},
		{colors.Black("Last Updated"), colors.DarkWhite(time.Unix(a.LastUpdated, 0).String())},
		// {colors.Black("Summary"), colors.DarkWhite(a.Summary)},
	})

	t.AddRow(colors.Black("Summary"), colors.DarkWhite(output.Fit(a.Summary, 62)))

	// t.SetAutoWrapText(true)
	// t.SetReflowDuringAutoWrap(true)

	var alertDetails string
	n := 0
	for k, v := range a.Details {
		line := fmt.Sprintf("%s: %s", k, v)
		if n == len(a.Details) {
			alertDetails = fmt.Sprintf("%s%s", alertDetails, line)
		} else {
			alertDetails = fmt.Sprintf("%s%s\n", alertDetails, line)
		}
		n++
	}
	t.AddRow(colors.Black("Details"), colors.DarkWhite(alertDetails))

	t.Render()

	// fmt.Println()

	if a.Comments == nil {
		return
	}

	output.SubTitleT2("Notes")

	alertComments(a)
}

func alertComments(a *events.Alert) {
	if a.Comments == nil {
		return
	}

	var tmSort []int64
	for tm := range a.Comments {
		tmSort = append(tmSort, tm)
	}
	sort.Slice(tmSort, func(i, j int) bool { return tmSort[i] < tmSort[j] })

	t := table.New()
	// t.SetRowLine("─")

	sep := strings.Repeat("─", 50)
	n := len(a.Comments)
	var m int

	for _, tm := range tmSort {
		c := a.Comments[tm]
		m++
		author := output.UserLocal(fmt.Sprintf("%s @%s", c.UserEmail, a.AccountID))
		timestamp := time.Unix(tm, 0).String()
		hdr := fmt.Sprintf("%s %s", colors.Black(timestamp), author)
		t.AddRow(hdr)
		t.AddRow(fmt.Sprintf("%s", c.Text))
		if m < n {
			t.AddRow(colors.Black(sep))
		}
	}

	t.Render()
	fmt.Println()
}

func (api *API) AlertStatus(s string) string {
	switch strings.ToLower(s) {
	case "triggered":
		return output.StrError(strings.ToUpper(s))
	case "acknowledged":
		return output.StrWarning(strings.ToUpper(s))
	case "resolved":
		return output.StrOk(strings.ToUpper(s))
	}

	return output.StrInactive(strings.ToUpper(s))
}

func alertSeverity(s string) string {
	switch strings.ToLower(s) {
	case "info":
		return output.StrBlue(strings.ToUpper(s))
	case "warning":
		return output.StrWarning(strings.ToUpper(s))
	case "error":
		return output.StrError(strings.ToUpper(s))
	case "critical":
		return output.StrMagenta(strings.ToUpper(s))
	}

	return output.StrInactive(strings.ToUpper(s))
}
