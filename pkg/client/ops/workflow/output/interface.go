package output

import "mmesh.dev/m-api-go/grpc/resources/ops/workflow"

type Interface interface {
	List(workflows map[string]*workflow.Workflow)
	Show(wf *workflow.Workflow)
}
type API struct{}
