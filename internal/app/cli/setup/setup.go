package setup

import (
	"fmt"

	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func Configure(configFile string) bool {
	output.Header("Setup")

	var isConfigured bool

	if input.GetConfirm("Want to configure the CLI for an existing account?", false) {
		fmt.Println()
		setupExistingAccount(configFile)
		return true
	}

	if input.GetConfirm("Want to create a new account?", false) {
		fmt.Println()
		client.Account().New()
		return true
	}

	fmt.Println()

	return isConfigured
}
