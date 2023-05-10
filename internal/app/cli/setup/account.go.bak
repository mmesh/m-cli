package setup

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/controller"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/config"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
)

func setupExistingAccount() {
	client.Auth().Login(login.NewRequest(), true)

	accountID := viper.GetString("logged.realm")
	userEmail := viper.GetString("logged.email")

	if len(accountID) == 0 {
		status.Error(fmt.Errorf("missing accountID"), "Invalid accountID")
	}

	if len(userEmail) == 0 {
		status.Error(fmt.Errorf("missing email"), "Invalid email")
	}

	viper.Set("account.id", accountID)

	vars.AccountID = accountID

	a := account.GetAccount()

	nxc, grpcConn := grpc.GetManagerProviderAPIClient(true)
	defer grpcConn.Close()

	a, err := nxc.GetAccount(context.TODO(), a)
	if err != nil {
		status.Error(err, "Unable to get account")
	}

	f := &controller.Federation{
		LocationID:   a.LocationID,
		FederationID: a.FederationID,
	}

	f, err = nxc.GetFederation(context.TODO(), f)
	if err != nil {
		status.Error(err, "Unable to get federation")
	}

	var cAuthServer, cEndpoint string
	for _, c := range f.Controllers {
		cAuthServer = fmt.Sprintf("https://%s", c.VirtualHost)
		cEndpoint = fmt.Sprintf("%s:%d", c.VirtualHost, c.Port)
		break
	}

	if err := config.DefaultAccount(cAuthServer, cEndpoint, accountID, userEmail); err != nil {
		status.Error(err, "Unable to write configuration file")
	}

	fmt.Printf("Ready to go :)\n\n")
}
