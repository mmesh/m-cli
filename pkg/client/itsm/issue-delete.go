package itsm

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *issueAPI) Delete() {
	i := getIssue(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	ir := &itsm.IssueRequest{
		AccountID: i.AccountID,
		IssueID:   i.IssueID,
	}

	sr, err := nxc.DeleteIssue(context.TODO(), ir)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete ITSM issue")
	}

	s.Stop()

	output.Show(sr)
}
