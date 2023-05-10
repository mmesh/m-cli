package tenant

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Update() {
	t := GetTenant()

	utr := &tenant.UpdateTenantRequest{
		AccountID:   t.AccountID,
		TenantID:    t.TenantID,
		Name:        input.GetInput("Tenant Name:", "", t.Name, survey.Required),
		Description: input.GetInput("Description:", "", t.Description, survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	t, err := nxc.UpdateTenant(context.TODO(), utr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set tenant")
	}

	s.Stop()

	// output.Show(t)
	Output().Show(t)
}
