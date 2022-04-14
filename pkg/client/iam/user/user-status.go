package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Enable() {
	setStatus(true)
}

func (api *API) Disable() {
	setStatus(false)
}

func setStatus(enabled bool) {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	ur := &iam.UserRequest{
		AccountID: u.AccountID,
		Email:     u.Email,
	}

	s := output.Spinner()

	var err error
	if enabled {
		u, err = nxc.EnableUser(context.TODO(), ur)
	} else {
		u, err = nxc.DisableUser(context.TODO(), ur)
	}
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user")
	}

	s.Stop()

	Output().Show(u)
}
