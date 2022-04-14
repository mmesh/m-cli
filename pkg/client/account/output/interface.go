package output

import (
	"mmesh.dev/m-api-go/grpc/resources/account"
	user_output "mmesh.dev/m-cli/pkg/client/iam/user/output"
	"mmesh.dev/m-cli/pkg/output"
)

type Interface interface {
	show(a *account.Account)

	Show(a *account.Account)

	Basic(a *account.Account)
	Admin(a *account.Account)
	Subscriptions(a *account.Account)
	Configuration(a *account.Account)
	Stats(a *account.Account)
}
type API struct{}

func (api *API) Show(a *account.Account) {
	api.show(a)
	showServiceSubscription(a)
	showIntegrations(a.Integrations)
	// showTraffic(a.Traffic)
	showUsage(a.Usage)
}

func (api *API) Basic(a *account.Account) {
	api.show(a)
}

func (api *API) Admin(a *account.Account) {
	// api.show(a)
	output.SectionHeader("Account Details")
	output.TitleT1("Account Admin Information")

	// output.SubTitleT1("Account Admin")
	user_output.ShowUser(a.Owner.Admin)
}

func (api *API) Subscriptions(a *account.Account) {
	api.show(a)
	showServiceSubscription(a)
}

func (api *API) Configuration(a *account.Account) {
	api.show(a)
	showIntegrations(a.Integrations)
}

func (api *API) Stats(a *account.Account) {
	api.show(a)
	// showTraffic(a.Traffic)
	showUsage(a.Usage)
}
