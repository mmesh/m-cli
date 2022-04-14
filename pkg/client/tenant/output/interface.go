package output

import "mmesh.dev/m-api-go/grpc/resources/tenant"

type Interface interface {
	List(tenants map[string]*tenant.Tenant)
	Show(t *tenant.Tenant)
}
type API struct{}
