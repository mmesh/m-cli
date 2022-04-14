package output

import "mmesh.dev/m-api-go/grpc/resources/ops/project"

type Interface interface {
	List(projects map[string]*project.Project)
	Show(p *project.Project)
}
type API struct{}
