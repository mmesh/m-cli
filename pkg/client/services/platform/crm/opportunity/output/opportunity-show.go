package output

import (
	"fmt"
	"strings"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(o *crm.Opportunity) {
	output.SectionHeader("CRM: Sales Opportunity Details")
	output.TitleT1("Sales Opportunity Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(o.AccountID))
	t.AddRow(colors.Black("Opportunity ID"), colors.DarkWhite(o.OpportunityID))

	if len(o.CustomerCompany) > 0 {
		t.AddRow(colors.Black("Company"), colors.DarkWhite(strings.Title(o.CustomerCompany)))
	}
	t.AddRow(colors.Black("Contact"), colors.DarkWhite(o.CustomerEmail))

	t.AddRow(colors.Black("Product/Service Name"), colors.DarkWhite(o.Spec.Product.Name))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(o.Spec.Product.Description))

	t.AddRow(colors.Black("Milestone"), milestone(o.Milestone))
	t.AddRow(colors.Black("Outcome"), outcome(o.Outcome))
	if o.Outcome == crm.OpportunityOutcome_OUTCOME_LOST {
		t.AddRow(colors.Black("Lost Reason"), outcomeLostReason(o.LostReason))
	}

	if o.LastContacted != 0 {
		tm := time.Unix(o.LastContacted, 0).String()
		t.AddRow(colors.Black("Last Contacted"), colors.DarkWhite(tm))
	}

	tm := time.Unix(o.CreationDate, 0).String()
	t.AddRow(colors.Black("Creation Date"), colors.DarkWhite(tm))
	tm = time.Unix(o.LastModified, 0).String()
	t.AddRow(colors.Black("Last Modified"), colors.DarkWhite(tm))

	t.Render()
	fmt.Println()

	if len(o.CustomerText) > 0 {
		output.SubTitleT2("Customer Requirements / Custom Details")
		// fmt.Println("Customer Requirements / Custom Details")
		// fmt.Println(colors.Black("======================================"))
		fmt.Println(colors.DarkWhite(o.CustomerText))

		fmt.Println()
	}
}

func milestone(m crm.OpportunityMilestone) string {
	switch m {
	case crm.OpportunityMilestone_NEW:
		return output.StrNormal("NEW")
	case crm.OpportunityMilestone_QUALIFIED:
		return output.StrNormal("QUALIFIED")
	case crm.OpportunityMilestone_MEETING:
		return output.StrNormal("MEETING")
	case crm.OpportunityMilestone_PROPOSAL:
		return output.StrNormal("PROPOSAL")
	case crm.OpportunityMilestone_NEGOTIATION:
		return output.StrOk("NEGOTIATION")
	case crm.OpportunityMilestone_CONTRACT:
		return output.StrOk("CONTRACT")
	}

	return output.StrInactive("-")
}

func outcome(o crm.OpportunityOutcome) string {
	switch o {
	case crm.OpportunityOutcome_OUTCOME_WON:
		return output.StrOk("WON")
	case crm.OpportunityOutcome_OUTCOME_LOST:
		return output.StrNormal("LOST")
	}

	return output.StrInactive("-")
}

func outcomeLostReason(lr crm.OpportunityLostReason) string {
	switch lr {
	case crm.OpportunityLostReason_LOST_ABANDONED:
		return output.StrError("ABANDONED")
	case crm.OpportunityLostReason_LOST_BECAME_STALE:
		return output.StrError("BECAME_STALE")
	case crm.OpportunityLostReason_LOST_COMPETITION:
		return output.StrError("COMPETITION")
	case crm.OpportunityLostReason_LOST_NO_AUTHORITY:
		return output.StrError("NO_AUTHORITY")
	case crm.OpportunityLostReason_LOST_POOR_QUALIFICATION:
		return output.StrError("POOR_QUALIFICATION")
	}

	return output.StrInactive("UNKNOWN")
}
