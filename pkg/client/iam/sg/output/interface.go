package output

import (
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/tenant"
)

type Interface interface {
	List(securityGroups map[string]*iam.SecurityGroup)
	Show(sg *iam.SecurityGroup, tenantMap map[string]*tenant.Tenant)
}
type API struct{}
