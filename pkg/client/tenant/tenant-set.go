package tenant

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	t := GetTenant(true)
	if t != nil { // editing existing resource
		output.Choice("Edit Tenant")
	} else { // <new> resource
		output.Choice("New Tenant")

		t = &tenant.Tenant{
			AccountID: a.AccountID,
		}

		if err := survey.AskOne(
			&survey.Input{Message: "Tenant ID:"},
			&t.TenantID,
			survey.WithValidator(validTenantID),
			survey.WithIcons(input.SurveySetIcons),
		); err != nil {
			status.Error(err, "Unable to get response")
		}
	}

	t.Description = input.GetInput("Description:", "", t.Description, survey.Required)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	t, err := nxc.SetTenant(context.TODO(), t)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set tenant")
	}

	s.Stop()

	// output.Show(t)
	Output().Show(t)
}
