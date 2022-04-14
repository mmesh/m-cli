package output

import "mmesh.dev/m-api-go/grpc/resources/ops/workflow"

type Interface interface {
	List(operations []*workflow.Operation)
	Show(o *workflow.Operation)
}
type API struct{}
