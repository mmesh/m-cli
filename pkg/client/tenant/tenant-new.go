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

func (api *API) New() {
	a := account.GetAccount()

	ntr := &tenant.NewTenantRequest{
		AccountID:   a.AccountID,
		Name:        input.GetInput("Tenant Name:", "", "", validTenantName),
		Description: input.GetInput("Description:", "", "", survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	t, err := nxc.CreateTenant(context.TODO(), ntr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set tenant")
	}

	s.Stop()

	// output.Show(t)
	Output().Show(t)
}
