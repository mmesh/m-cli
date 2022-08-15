package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func showService(a *account.Account) {
	if a.Service == nil {
		return
	}

	svc := a.Service

	output.SubTitleT2("Service Subscriptions")

	t := table.New()

	var svcStatus, canceled string
	if svc.Suspended {
		svcStatus = output.StrError("suspended")
	} else {
		svcStatus = output.StrOk("active")
	}
	if svc.Canceled {
		canceled = output.StrError("canceled")
	}

	serviceStatus := fmt.Sprintf("%s %s", svcStatus, canceled)

	t.AddRow(colors.Black("Service Status"), serviceStatus)

	// if a.Service.Traffic != nil {
	// 	if a.Service.Traffic.EnforcedRestriction {
	// 		t.AddRow("Managed Relay Service:", output.StrDisabled("restricted"))
	// 	} else {
	// 		t.AddRow("Managed Relay Service:", output.StrEnabled("active"))
	// 	}
	// }

	if a.Owner.Customer != nil {
		// if a.Owner.Customer.Delinquent {
		// 	t.AddRow(colors.Black("Latest Charge Status"), output.StrError("failed"))
		// }

		dpm := false
		if a.Owner.Customer.StripeDefaultPaymentMethod != nil {
			if len(a.Owner.Customer.StripeDefaultPaymentMethod.ID) > 0 {
				dpm = true
			}
		}
		if dpm {
			t.AddRow(colors.Black("Default Payment Method"), output.StrEnabled("active"))
		} else {
			t.AddRow(colors.Black("Default Payment Method"), output.StrDisabled("not-set"))
		}

		// ps := fmt.Sprintf("%d", len(a.Owner.Customer.StripePaymentSources))
		// t.AddRow(colors.Black("Saved Payment Sources"), colors.DarkWhite(ps))
		t.AddRow(colors.Black("Customer Balance"),
			output.AmountMoney(a.Owner.Customer.Balance, a.Owner.Customer.Currency))
		// output.CustomerBalance(a.Owner.Customer.Balance, a.Owner.Customer.Currency))
	}

	t.Render()
	fmt.Println()
}
