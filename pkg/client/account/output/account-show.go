package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) show(a *account.Account) {
	output.SectionHeader("Account Details")
	output.TitleT1("Account Information")

	t := table.New()

	var testing string
	if !a.Owner.Customer.LiveMode {
		testing = output.StrNormal("test")
	}

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(a.AccountID)+" "+testing)

	if len(a.Name) > 0 {
		t.AddRow(colors.Black("Name"), colors.DarkWhite(a.Name))
	}
	if len(a.Description) > 0 {
		t.AddRow(colors.Black("Description"), colors.DarkWhite(a.Description))
	}
	if len(a.CompanyName) > 0 {
		t.AddRow(colors.Black("Company"), colors.DarkWhite(a.CompanyName))
	}

	t.AddRow(colors.Black("Location ID"), colors.DarkWhite(a.LocationID))
	t.AddRow(colors.Black("Federation ID"), colors.DarkWhite(a.FederationID))
	// t.AddRow(colors.Black("Account Admin"), colors.DarkWhite(a.Owner.Admin.UserID))

	// if a.Type != account.AccountType_NORMAL {
	// 	t.AddRow(colors.Black("Account Type"), output.StrEnabled("provider"))
	// }

	if a.CreationDate > 0 {
		tm := time.Unix(a.CreationDate, 0)
		t.AddRow(colors.Black("Creation Date"), colors.DarkWhite(tm.String()))
	}
	var confirmed, status string
	if a.Status != nil {
		if a.Status.Confirmed {
			confirmed = output.StrOk("confirmed")
		} else {
			confirmed = output.StrWarning("not-confirmed")
		}
		if a.Status.Enabled {
			status = output.StrEnabled("enabled")
		} else {
			status = output.StrDisabled("disabled")
		}
	} else {
		status = output.StrInactive("not-activated")
	}
	var overLimit string
	if a.Usage != nil {
		if a.Usage.Limit != nil {
			if a.Usage.Limit.OverLimit {
				overLimit = output.StrWarning("over-limit")
			}
		}
	}
	t.AddRow(colors.Black("Account Status"), confirmed+" "+status+" "+overLimit)

	t.Render()
	fmt.Println()

	// showServiceSubscription(a)
	// showIntegrations(a.Cfg)
	// showTraffic(a.Traffic)
	// showUsage(a.Usage)
}
