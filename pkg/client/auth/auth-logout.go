package auth

import (
	"context"
	"os"

	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Logout() {
	s := output.Spinner()

	nxc, grpcConn := grpc.GetManagerAPIClient(true)
	defer grpcConn.Close()

	accountID, err := auth.GetAccountID()
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to find auth accountID")
	}

	userID, err := auth.GetUserID()
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to find auth userID")
	}

	req := &auth_pb.SignoutRequest{
		AccountID: accountID,
		UserID:    userID,
		// SessionID:
	}

	resp, err := nxc.Signout(context.TODO(), req)
	if err != nil {
		grpcConn.Close()
		s.Stop()
		status.Error(err, "Unable to signout")
	}

	if resp.Result != auth_pb.SignoutResult_SIGNOUT_SUCCESSFUL {
		grpcConn.Close()
		s.Stop()
		status.Error(err, "Unable to signout")
	}

	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to find API key")
	}

	if err := os.RemoveAll(apiKeyFile); err != nil {
		s.Stop()
		status.Error(err, "Unable to logout")
	}

	s.Stop()

	output.Logout()
}
