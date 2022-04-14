package tenant

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var tenantsMap map[string]*tenant.Tenant = nil
var selectedTenant *tenant.Tenant = nil

func GetTenant(edit bool) *tenant.Tenant {
	if selectedTenant != nil {
		return selectedTenant
	}

	tl := Tenants()

	if len(tl) == 0 {
		msg.Info("No tenant found")
		os.Exit(1)
	}

	var tenantOptID string
	tenantsOpts := make([]string, 0)
	tenants := make(map[string]*tenant.Tenant)

	for _, t := range tl {
		tenantOptID = t.TenantID
		tenantsOpts = append(tenantsOpts, tenantOptID)
		tenants[tenantOptID] = t
	}

	sort.Strings(tenantsOpts)

	if edit {
		tenantsOpts = append(tenantsOpts, input.NewResource)
	}

	tenantOptID = input.GetSelect("Tenant:", "", tenantsOpts, survey.Required)

	if tenantOptID == input.NewResource {
		return nil
	}

	vars.TenantID = tenants[tenantOptID].TenantID
	selectedTenant = tenants[tenantOptID]

	return tenants[tenantOptID]
}

func Tenants() map[string]*tenant.Tenant {
	if tenantsMap != nil {
		return tenantsMap
	}

	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &tenant.ListTenantsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	tenants := make(map[string]*tenant.Tenant)

	for {
		tl, err := nxc.ListTenants(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list tenants")
		}

		for _, t := range tl.Tenants {
			tenants[t.TenantID] = t
		}

		if len(tl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = tl.Meta.NextPageToken
		} else {
			break
		}
	}

	tenantsMap = tenants

	return tenants
}

func validTenantID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	tenantID := val.(string)

	if tenantsMap == nil {
		tenantsMap = Tenants()
	}

	if _, ok := tenantsMap[tenantID]; ok {
		return fmt.Errorf("tenant %s already exist", tenantID)
	}

	return nil
}
