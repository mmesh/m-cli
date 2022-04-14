package output

import (
	"fmt"
	"strings"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(i *cloud.Instance) {
	output.SectionHeader("Cloud: Instance Details")
	output.TitleT1("Instance Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(i.AccountID))
	t.AddRow(colors.Black("Instance ID"), colors.DarkWhite(i.InstanceID))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(i.Description))
	t.AddRow(colors.Black("Provider"), output.StrNormal(strings.Title(i.ProviderName)))
	t.AddRow(colors.Black("Region"), colors.DarkWhite(i.Spec.Location.Name))

	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(i.NodeInstance.Node.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(i.NodeInstance.Node.NetID))
	t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(i.NodeInstance.Node.VRFID))
	t.AddRow(colors.Black("Node ID"), colors.DarkWhite(i.NodeInstance.Node.NodeID))

	t.AddRow(colors.Black("Image"), colors.DarkWhite(i.Spec.Image.ImageID))
	cpu := fmt.Sprintf("%d", i.Spec.InstanceType.VCPUs)
	t.AddRow(colors.Black("VCPUs"), colors.DarkWhite(cpu))
	mem := fmt.Sprintf("%d MB", i.Spec.InstanceType.Memory)
	t.AddRow(colors.Black("RAM"), colors.DarkWhite(mem))

	if i.Spec.InstanceType.Disk > 0 {
		disk := fmt.Sprintf("%d GB", i.Spec.InstanceType.Disk)
		t.AddRow(colors.Black("Disk"), colors.DarkWhite(disk))
	}

	tm := time.Unix(i.CreationDate, 0).String()
	t.AddRow(colors.Black("Creation Date"), colors.DarkWhite(tm))
	tm = time.Unix(i.LastModified, 0).String()
	t.AddRow(colors.Black("Last Modified"), colors.DarkWhite(tm))

	t.Render()
	fmt.Println()

	if i.LastAction == nil && len(i.Actions) == 0 {
		return
	}

	output.SubTitleT2("Instance History")

	t = table.New()

	j := 0
	if i.LastAction != nil {
		j++

		var actionID uint64
		if i.LastAction.ActionID == 0 {
			actionID = uint64(j)
		} else {
			actionID = i.LastAction.ActionID
		}

		aID := fmt.Sprintf("%s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%d", j)), colors.DarkMagenta("]"), colors.Black("Action ID"))
		t.AddRow(aID, colors.DarkWhite(fmt.Sprintf("%d", actionID)))
		t.AddRow(colors.Black("    Action Type"), colors.White(strings.ToUpper(i.LastAction.Type)))
		t.AddRow(colors.Black("    Status"), actionStatus(i.LastAction.Status))

		if i.LastAction.StartDate != 0 {
			tm := time.Unix(i.LastAction.StartDate, 0).String()
			t.AddRow(colors.Black("    Started"), colors.DarkWhite(tm))
		}
		if i.LastAction.CompletionDate != 0 {
			tm := time.Unix(i.LastAction.CompletionDate, 0).String()
			t.AddRow(colors.Black("    Completed"), colors.DarkWhite(tm))
		}
		t.AddRow("", "")
	}

	for _, a := range i.Actions {
		j++

		var actionID uint64
		if a.ActionID == 0 {
			actionID = uint64(j)
		} else {
			actionID = a.ActionID
		}

		aID := fmt.Sprintf("%s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%d", j)), colors.DarkMagenta("]"), colors.Black("Action ID"))
		t.AddRow(aID, colors.DarkWhite(fmt.Sprintf("%d", actionID)))
		t.AddRow(colors.Black("    Action Type"), colors.White(strings.ToUpper(a.Type)))
		t.AddRow(colors.Black("    Status"), actionStatus(a.Status))

		if a.StartDate != 0 {
			tm := time.Unix(a.StartDate, 0).String()
			t.AddRow(colors.Black("    Started"), colors.DarkWhite(tm))
		}
		if a.CompletionDate != 0 {
			tm := time.Unix(a.CompletionDate, 0).String()
			t.AddRow(colors.Black("    Completed"), colors.DarkWhite(tm))
		}
		t.AddRow("", "")
	}

	t.Render()
	fmt.Println()
}

func actionStatus(s string) string {
	switch s {
	case "in-progress":
		return output.StrNormal(s)
	case "completed":
		return output.StrOk(s)
	case "errored":
		return output.StrError(s)
	}

	return output.StrNormal(s)
}
