package output

import "mmesh.dev/m-api-go/grpc/resources/iam"

type Interface interface {
	List(roles map[string]*iam.Role)
	Show(r *iam.Role)
}
type API struct{}
