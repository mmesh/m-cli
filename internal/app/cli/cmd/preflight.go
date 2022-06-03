package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
)

func header() {
	fmt.Println(colors.Black(version.CLI_NAME + " " + version.GetVersion()))
	output.AppHeader(vars.AccountID, false)
}

func appHeader(str string) string {
	h1 := colors.Black(version.CLI_NAME + " " + version.GetVersion())
	h2 := output.AppHeader(vars.AccountID, true)

	return fmt.Sprintf("%s\n%s%s", h1, h2, str)
}

func preflightNoLogin() {
	header()

	if !isConfigured {
		notConfigured()
		os.Exit(0)
	}

	accountID := viper.GetString("account.id")
	if len(accountID) == 0 {
		status.Error(fmt.Errorf("missing account.id"), "Unable to get configured account")
	}
}

func preflight() {
	header()

	if isConfigured {
		// silent auto-login
		autoLogin()
	} else {
		notConfigured()
		os.Exit(0)
	}
}

func autoLogin() {
	accountID := viper.GetString("account.id")
	if len(accountID) == 0 {
		status.Error(fmt.Errorf("missing account.id"), "Unable to get configured account")
	}

	if client.Auth().LoginRequired(accountID) {
		client.Auth().Login(login.NewRequest(), true)
	}
	// client.Auth().AutoLogin(login.NewRequest())
}

func notConfigured() {
	msg.Error("Configuration not detected")

	fmt.Printf("%s\n", colors.DarkBlue("_"))
	cmd := colors.DarkWhite(fmt.Sprintf("%s setup", version.CLI_NAME))
	q := colors.DarkBlue("'")
	msg := fmt.Sprintf("%s %s%s%s", colors.Black("Please configure the client with"), q, cmd, q)
	fmt.Printf("%s %s\n\n", colors.Cyan("🢂"), msg)
}
