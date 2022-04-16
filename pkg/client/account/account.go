package account

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
)

func GetAccount() *account.Account {
	accountID := vars.AccountID

	if len(accountID) == 0 {
		accountID = viper.GetString("account.id")
		if len(accountID) == 0 {
			accountID = input.GetInput("Account ID:", "", "", survey.Required)
		}
		vars.AccountID = accountID
	}

	return &account.Account{
		AccountID:   accountID,
		Description: "Account " + accountID,
	}
}

func FetchAccount() *account.Account {
	a := GetAccount()

	nxc, grpcConn := grpc.GetControllerProviderAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	a, err := nxc.GetAccount(context.TODO(), a)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get account")
	}

	s.Stop()

	return a
}

func fetchAccountStats() *account.Account {
	a := GetAccount()

	nxc, grpcConn := grpc.GetControllerProviderAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	if _, err := nxc.GetAccountStats(context.TODO(), a); err != nil {
		s.Stop()
		status.Error(err, "Unable to get account stats")
	}

	a, err := nxc.GetAccount(context.TODO(), a)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get account")
	}

	s.Stop()

	return a
}
