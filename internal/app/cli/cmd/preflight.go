package cmd

import (
	"fmt"
	"os"

	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/output"
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

func preflight() {
	header()

	if isConfigured {
		// silent auto-login
		autoLogin()
	} else {
		msg.Error("Configuration not detected")

		fmt.Printf("%s\n", colors.DarkBlue("_"))
		cmd := colors.DarkWhite(fmt.Sprintf("%s setup", version.CLI_NAME))
		q := colors.DarkBlue("'")
		msg := fmt.Sprintf("%s %s%s%s", colors.Black("Please configure the client with"), q, cmd, q)
		fmt.Printf("%s %s\n\n", colors.Cyan("🢂"), msg)

		os.Exit(0)
	}
}

func autoLogin() {
	if len(os.Args) > 1 {
		if os.Args[1] == "auth" {
			return
		}
	}

	if client.Auth().LoginRequired() {
		client.Auth().Login(login.NewRequest(), true)
	}
	// client.Auth().AutoLogin(login.NewRequest())
}
