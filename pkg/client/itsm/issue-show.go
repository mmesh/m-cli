package itsm

import "mmesh.dev/m-api-go/grpc/resources/itsm"

func (api *issueAPI) Show() *itsm.Issue {
	i := getIssue(false)

	Output().Show(i)

	return i
}
