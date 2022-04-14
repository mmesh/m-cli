package output

import (
	"fmt"
	"strings"
	"time"

	"github.com/skratchdot/open-golang/open"
	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func showServiceSubscription(a *account.Account) {
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
		if a.Owner.Customer.Delinquent {
			t.AddRow(colors.Black("Latest Charge Status"), output.StrError("failed"))
		}
		pm := fmt.Sprintf("%d", len(a.Owner.Customer.StripePaymentMethods))
		t.AddRow(colors.Black("Saved Payment Methods"), colors.DarkWhite(pm))
		t.AddRow(colors.Black("Customer Balance"), output.CustomerBalance(a.Owner.Customer.Balance, a.Owner.Customer.Currency))
	}

	t.Render()
	fmt.Println()

	n := 0
	for _, s := range svc.Subscriptions {
		t := table.New()

		svcID := fmt.Sprintf("%s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%02d", n+1)), colors.DarkMagenta("]"), colors.Black("Service ID"))
		t.AddRow(svcID, colors.DarkWhite(s.ServiceID))
		if len(s.PricingPlanID) > 0 {
			t.AddRow(colors.Black("     Plan ID"), output.StrGreen(strings.ToUpper(s.PricingPlanID)))
		}
		// if len(s.ProviderID) > 0 {
		// 	t.AddRow(colors.Black("     Provider ID"), colors.DarkWhite(strings.Title(s.ProviderID)))
		// }
		if len(s.ProductID) > 0 {
			t.AddRow(colors.Black("     Product ID"), colors.DarkWhite(strings.Title(s.ProductID)))
		}
		if len(s.PriceID) > 0 {
			t.AddRow(colors.Black("     Price ID"), colors.DarkWhite(strings.Title(s.PriceID)))
		}

		if s.Discount != nil {
			if len(s.Discount.PercentOff) > 0 {
				percentOff := fmt.Sprintf("%s%%", s.Discount.PercentOff)
				t.AddRow(colors.Black("     Promotional Discount"), colors.Green(percentOff))
			}
		}

		var openInvoice bool
		var sstatus string
		switch s.Status {
		case billing.SubscriptionStatus_UNKNOWN:
			sstatus = output.StrInactive("unknown")
		case billing.SubscriptionStatus_NONE:
			sstatus = output.StrInactive("none")
		case billing.SubscriptionStatus_TRIALING:
			sstatus = output.StrEnabled("trial")
		case billing.SubscriptionStatus_ACTIVE:
			sstatus = output.StrEnabled("active")
		case billing.SubscriptionStatus_INCOMPLETE:
			openInvoice = true
			sstatus = output.StrWarning("incomplete")
		case billing.SubscriptionStatus_INCOMPLETE_EXPIRED:
			openInvoice = true
			sstatus = output.StrError("incomplete-expired")
		case billing.SubscriptionStatus_PAST_DUE:
			openInvoice = true
			sstatus = output.StrWarning("past-due")
		case billing.SubscriptionStatus_CANCELED:
			openInvoice = true
			sstatus = output.StrError("canceled")
		case billing.SubscriptionStatus_UNPAID:
			openInvoice = true
			sstatus = output.StrError("unpaid")
		default:
			sstatus = output.StrInactive("n/a")
		}

		t.AddRow(colors.Black("     Status"), sstatus)

		tm := time.Unix(s.StartDate, 0).String()
		t.AddRow(colors.Black("     Start Date"), colors.DarkWhite(tm))
		if s.EndDate != 0 {
			tm = time.Unix(s.EndDate, 0).String()
			t.AddRow(colors.Black("     End Date"), colors.DarkWhite(tm))
		}

		if s.TrialStartDate != 0 {
			tm = time.Unix(s.TrialStartDate, 0).String()
			t.AddRow(colors.Black("     Trial Period Start"), colors.DarkWhite(tm))
			tm = time.Unix(s.TrialEndDate, 0).String()
			t.AddRow(colors.Black("     Trial Period End"), colors.DarkWhite(tm))
		}

		tm = time.Unix(s.CurrentPeriodStart, 0).String()
		t.AddRow(colors.Black("     Current Period Start"), colors.DarkWhite(tm))
		tm = time.Unix(s.CurrentPeriodEnd, 0).String()
		t.AddRow(colors.Black("     Current Period End"), colors.DarkWhite(tm))

		if s.CancelAtPeriodEnd {
			t.AddRow(colors.Black("     Cancel at Period End"), output.StrWarning("yes"))
			tm = time.Unix(s.CancelationDate, 0).String()
			t.AddRow(colors.Black("     Cancelation Date"), colors.DarkWhite(tm))
		}

		tm = time.Unix(s.CreationDate, 0).String()
		t.AddRow(colors.Black("     Creation Date"), colors.DarkWhite(tm))

		tm = time.Unix(s.LastModified, 0).String()
		t.AddRow(colors.Black("     Last Modified"), colors.DarkWhite(tm))

		t.Render()
		fmt.Println()

		if openInvoice && len(s.LatestStripeHostedInvoiceURL) > 0 {
			if input.GetConfirm("Detected unpaid invoice, open payment form now?", true) {
				if err := open.Start(s.LatestStripeHostedInvoiceURL); err != nil {
					status.Error(err, "Unable to open URL in your browser")
				}

				fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening URL in your browser..."))
			} else {
				fmt.Println()
			}
		}

		n++
	}
}
