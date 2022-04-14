package setup

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func Setup(configFile string) error {
	output.Header("First-time Setup")

	var existingAccount bool
	prompt := &survey.Confirm{Message: "Want to configure the CLI for an existing account?"}
	if err := survey.AskOne(prompt, &existingAccount, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	fmt.Println()

	if existingAccount {
		setupExistingAccount(configFile)
	} else {
		// setupNewAccount()
		client.Account().New()
	}

	// return viper.WriteConfigAs(configFile)
	return nil
}
