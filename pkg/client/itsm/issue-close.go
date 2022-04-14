package itsm

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *issueAPI) Close() {
	i := getIssue(true)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	ir := &itsm.IssueRequest{
		AccountID: i.AccountID,
		IssueID:   i.IssueID,
	}

	sr, err := nxc.CloseIssue(context.TODO(), ir)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to close ITSM issue")
	}

	s.Stop()

	output.Show(sr)
}
