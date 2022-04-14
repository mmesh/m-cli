package output

import "mmesh.dev/m-api-go/grpc/resources/iam"

type Interface interface {
	List(users map[string]*iam.User)
	Show(u *iam.User)
}
type API struct{}
