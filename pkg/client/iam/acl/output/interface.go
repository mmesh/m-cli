package output

import (
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

type Interface interface {
	List(acls map[string]*iam.ACL)
	Show(acl *iam.ACL, nsm *topology.NodeSummaryMap)
}
type API struct{}
