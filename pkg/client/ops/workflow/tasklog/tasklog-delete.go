package tasklog

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	tl := GetTaskLog()

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetOpsAPIClient()
	defer grpcConn.Close()

	tlr := &ops.TaskLogReq{
		AccountID:  tl.AccountID,
		TenantID:   tl.TenantID,
		ProjectID:  tl.ProjectID,
		WorkflowID: tl.WorkflowID,
		TaskLogID:  tl.TaskLogID,
	}

	sr, err := nxc.DeleteTaskLog(context.TODO(), tlr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete operation")
	}

	s.Stop()

	output.Show(sr)
}
