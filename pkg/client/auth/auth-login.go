package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Login(req *auth_pb.LoginRequest, verbose bool) {
	nxc, grpcConn := grpc.GetManagerAPIClient(false)
	defer grpcConn.Close()

	resp, err := nxc.Login(context.TODO(), req)
	if err != nil {
		grpcConn.Close()
		status.Error(err, "Unable to login")
	}

	switch resp.Result {
	case auth_pb.LoginResult_LOGIN_FAILED:
		msg.Error("Login failed")
	case auth_pb.LoginResult_IAM_ACCOUNT_UNCONFIRMED:
		// unconfirmedAccount()
		msg.Error("Account not confirmed")
	case auth_pb.LoginResult_IAM_ACCOUNT_DISABLED:
		msg.Error("Account disabled, please contact customer service")
	case auth_pb.LoginResult_IAM_USER_UNCONFIRMED:
		// unconfirmedUser()
		msg.Error("User not confirmed")
	case auth_pb.LoginResult_IAM_USER_DISABLED:
		msg.Error("User disabled, please contact your mmesh account administrator")
	}

	if resp.Result != auth_pb.LoginResult_LOGIN_SUCCESSFUL {
		os.Exit(1)
	}

	ac := &auth.Credentials{
		AccountID:    resp.AccountID,
		FederationID: resp.FederationID,
		Key:          resp.AuthKey.Key,
	}

	if err := setAPIKey(ac); err != nil {
		status.Error(err, "Unable to set apiKey")
	}

	viper.Set("user.accountID", resp.AccountID)
	viper.Set("user.federationID", resp.FederationID)
	viper.Set("user.isAdmin", resp.IsAdmin)

	if verbose {
		fmt.Println()
		output.Authenticated()
	}
}

func setAPIKey(ac *auth.Credentials) error {
	jsonData, err := json.Marshal(ac)
	if err != nil {
		return errors.Wrapf(err, "[%v] function json.Marshal()", errors.Trace())
	}

	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		return errors.Wrapf(err, "[%v] function getAPIKeyFile()", errors.Trace())
	}

	if err := ioutil.WriteFile(apiKeyFile, jsonData, 0600); err != nil {
		return errors.Wrapf(err, "[%v] function ioutil.WriteFile()", errors.Trace())
	}

	return nil
}
