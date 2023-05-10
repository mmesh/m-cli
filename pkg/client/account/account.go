package account

import (
	"context"
	"os"

	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetAccount() *account.Account {
	var err error

	accountID := vars.AccountID

	if len(accountID) == 0 {
		accountID, err = auth.GetAccountID()
		if err != nil {
			msg.Alert("Unable to get accountID.")
			msg.Alert("Invalid or inexistent api key. Login to refresh your token.")
			os.Exit(1)
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

	nxc, grpcConn := grpc.GetAccountAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	ar := &account.AccountReq{
		AccountID: a.AccountID,
	}

	a, err := nxc.GetAccount(context.TODO(), ar)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get account")
	}

	s.Stop()

	return a
}

func fetchAccountStats() *account.Account {
	a := GetAccount()

	nxc1, grpcConn1 := grpc.GetControllerAPIClient()
	defer grpcConn1.Close()

	s := output.Spinner()

	ar := &account.AccountReq{
		AccountID: a.AccountID,
	}

	if _, err := nxc1.GetAccountStats(context.TODO(), ar); err != nil {
		s.Stop()
		status.Error(err, "Unable to get account stats")
	}

	nxc2, grpcConn2 := grpc.GetAccountAPIClient(true)
	defer grpcConn2.Close()

	a, err := nxc2.GetAccount(context.TODO(), ar)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get account")
	}

	s.Stop()

	return a
}
