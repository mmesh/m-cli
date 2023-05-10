package sg

import "mmesh.dev/m-cli/pkg/client/tenant"

func (api *API) Show() {
	tenantMap := tenant.Tenants()

	Output().Show(GetSecurityGroup(false), tenantMap)
}
