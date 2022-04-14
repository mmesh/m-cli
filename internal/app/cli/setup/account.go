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

func setupExistingAccount(configFile string) {
	client.Auth().Login(login.NewRequest(), true)

	accountID := viper.GetString("logged.realm")
	userEmail := viper.GetString("logged.email")
	if len(accountID) > 0 && len(userEmail) > 0 {
		vars.AccountID = accountID
		viper.Set("account.id", accountID)
		viper.Set("user.email", userEmail)
	}

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

	var controllerAuthServer, controllerEndpoint string
	for _, c := range f.Controllers {
		controllerAuthServer = fmt.Sprintf("https://%s", c.Host)
		controllerEndpoint = fmt.Sprintf("%s:%d", c.Host, c.Port)
		break
	}

	viper.Set("controller.authServer", controllerAuthServer)
	viper.Set("controller.endpoint", controllerEndpoint)

	fmt.Println()

	config.WriteCLIConfig(configFile)

	fmt.Printf("Ready to go :-)\n\n")
}
