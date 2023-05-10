package tenant

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	t := GetTenant()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetTenantAPIClient()
	defer grpcConn.Close()

	tr := &tenant.TenantReq{
		AccountID: t.AccountID,
		TenantID:  t.TenantID,
	}

	sr, err := nxc.DeleteTenant(context.TODO(), tr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete tenant")
	}

	s.Stop()

	output.Show(sr)
}
