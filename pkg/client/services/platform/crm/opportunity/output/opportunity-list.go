package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) List(opportunities map[string]*crm.Opportunity) {
	output.SectionHeader("CRM: Sales Opportunities")
	output.TitleT1("Sales Opportunity List")

	if len(opportunities) == 0 {
		msg.Info("No records found")
		return
	}

	t := table.New()

	t.SetColWidth(36)

	t.Header(colors.Black("CONTACT"), colors.Black("PRODUCT"))

	t.SetAutoWrapText(true)
	t.SetReflowDuringAutoWrap(true)
	t.SetRowLine("─")

	for _, o := range opportunities {
		var contact, company string
		if len(o.CustomerCompany) > 0 {
			company = fmt.Sprintf(" @%s", o.CustomerCompany)
		}
		if len(o.CustomerNickname) > 0 {
			contact = colors.DarkWhite(fmt.Sprintf("%s%s", o.CustomerNickname, company))
		} else {
			contact = colors.DarkWhite(fmt.Sprintf("%s%s", o.CustomerEmail, company))
		}
		t.AddRow(contact, o.Spec.Product.Name)
	}

	t.Render()
	fmt.Println()
}
