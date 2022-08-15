package account

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/biter777/countries"
	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/client/iam/user/credentials"
	"mmesh.dev/m-cli/pkg/config"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) New() {
	f := selectFederation()

	req := &account.NewAccountRequest{
		LocationID:   f.LocationID,
		FederationID: f.FederationID,
		Address:      &billing.Address{},
	}

	nxc, grpcConn := grpc.GetManagerProviderAPIClient(false)
	defer grpcConn.Close()

	req.AccountID = input.GetInput("Account ID:", "", "", survey.Required)
	req.Email = input.GetInput("Admin Email:", "", "", input.ValidEmail)
	req.Password = credentials.SetCredentialsPassword(true)

	// pp := pricingplan.GetPricingPlan()

	// var priceID string
	// var priceOpts []string
	// prices := make(map[string]*services.Price)
	// for _, p := range pp.Prices {
	// 	// priceID = fmt.Sprintf("%s %s/%s", strings.Title(p.PriceID), output.AmountMoney(p.UnitAmount, p.Currency), pp.UnitLabel)
	// 	priceID = fmt.Sprintf("%s %s (15-day free trial)", strings.Title(p.PriceID), output.AmountMoney(p.UnitAmount, p.Currency))
	// 	priceOpts = append(priceOpts, priceID)
	// 	prices[priceID] = p
	// }
	// if len(prices) > 0 {
	// 	priceID = input.GetSelect("Price:", "", priceOpts, survey.Required)
	// }

	// nUsers := input.GetInputInt("Users:", "Number of users", "1", input.ValidUint)

	// if pp.Type != services.PlanType_NETWORK {
	// 	status.Error(fmt.Errorf("subscription plan type not supported"), "Unable to create account")
	// }

	/*
		// paymentRequired := false
		if prices[priceID].UnitAmount > 0 || prices[priceID].UnitAmountDecimal > 0 {
			// paymentRequired = true
			fmt.Println()
			output.Header("Billing Information")
			req.Name = input.GetInput("Full Name:", "", "", survey.Required)
			req.Address = getAddress()
		} else {
			req.Address.Country = getCountry()
		}
	*/

	fmt.Println()
	output.Header("Billing Information")
	req.Name = input.GetInput("Full Name:", "", "", survey.Required)
	req.Address = getAddress()

	// req.PlanType = pp.Type
	req.PlanType = services.PlanType_NETWORK
	req.Subscription = &billing.Subscription{
		// StripePriceID: prices[priceID].StripePriceID,
		// Quantity:      int64(nUsers),
		Quantity: int64(1),
	}

	s := output.Spinner()

	if _, err := nxc.NewAccount(context.TODO(), req); err != nil {
		s.Stop()
		status.Error(err, "Unable to create account")
	}

	s.Stop()

	vars.AccountID = req.AccountID

	cAuthServer := fmt.Sprintf("https://%s", f.VirtualHost)
	cEndpoint := fmt.Sprintf("%s:%d", f.VirtualHost, f.Port)

	fmt.Println()

	if err := config.DefaultAccount(cAuthServer, cEndpoint, req.AccountID, req.Email); err != nil {
		status.Error(err, "Unable to write configuration")
	}

	welcomeMsg(req.AccountID, req.Email)

	/*
		if paymentRequired {
			if len(resp.HostedInvoiceURL) == 0 {
				status.Error(fmt.Errorf("missing hostedInvoiceURL"), "Unable to setup account billing")
			}

			if err := open.Start(resp.HostedInvoiceURL); err != nil {
				status.Error(err, "Unable to open URL in your browser")
			}

			fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening URL in your browser..."))
			// fmt.Printf("\n%s %s\n\n", colors.DarkWhite("->"), colors.DarkCyan(r.URL))
		}
	*/
}

func welcomeMsg(accountID, email string) {
	fmt.Printf(`
Congratulations and thanks for setting up your mmesh account!

A confirmation email has been sent to %s.

For convenience, your account %s was created with some defaults
and will be enabled and fully functional once we get your confirmation.

Follow the instructions you will find in the email and please notice
that if your confirmation is not received within 24h, the account will be
deleted.

If for any reason you don't get the confirmation email, you can execute
the command %s and the confirmation
instructions will be resent.

Once the account is active, login with '%s' and
use '%s' to create your first network,
then use '%s' to create the first subnet :-)

Now, depending on your subscription choice, you will be redirected in your
browser to our billing secure page. There you will be able to enter
your payment information to unleash all the powers of mmesh.

Then you will be able to provision and manage, right from this CLI and
seamlessly integrated in your mmesh, cloud resources (kubernetes clusters,
VM instances, ready-to-use apps, etc) and mmesh-certified managed services
from top-rated providers at the best prices.

Enjoy and welcome to your mmesh!
-mmesh team :-)

`,
		colors.White(email),
		colors.White(accountID),
		colors.White("mmeshctl auth resend-confirmation"),
		colors.White("mmeshctl auth login"),
		colors.White("mmeshctl network create"),
		colors.White("mmeshctl subnet create"))
}

func getAddress() *billing.Address {
	return &billing.Address{
		Line1:      input.GetInput("Address Line1:", "", "", survey.Required),
		Line2:      input.GetInput("Address Line2:", "", "", nil),
		PostalCode: input.GetInput("Postal Code:", "", "", survey.Required),
		City:       input.GetInput("City:", "", "", survey.Required),
		State:      input.GetInput("State:", "", "", survey.Required),
		Country:    getCountry(),
	}
}

func getCountry() string {
	var countryName string
	var countryOpts []string
	cl := make(map[string]*countries.Country)
	for _, c := range countries.AllInfo() {
		countryName = c.Name
		countryOpts = append(countryOpts, countryName)
		cl[c.Name] = c
	}

	countryName = input.GetSelect("Country:", "", countryOpts, survey.Required)

	return cl[countryName].Alpha2
}
