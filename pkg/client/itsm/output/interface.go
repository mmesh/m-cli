package output

import (
	"io"

	"mmesh.dev/m-api-go/grpc/resources/itsm"
)

type Interface interface {
	List(issues map[string]*itsm.Issue)
	Show(i *itsm.Issue)
	ChatIssueInfo(w io.Writer, i *itsm.Issue)
}
type API struct{}
