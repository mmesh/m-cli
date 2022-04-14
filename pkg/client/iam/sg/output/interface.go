package output

import "mmesh.dev/m-api-go/grpc/resources/iam"

type Interface interface {
	List(securityGroups map[string]*iam.SecurityGroup)
	Show(sg *iam.SecurityGroup)
}
type API struct{}
