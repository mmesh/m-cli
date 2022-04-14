package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) GitHub(gh *thirdParty.GitHub) *thirdParty.GitHub {
	output.Header("Github Setup")

	if gh == nil {
		gh = &thirdParty.GitHub{}
	}

	if gh.Enabled {
		if input.GetConfirm("Delete GitHub integration?", false) {
			gh = &thirdParty.GitHub{}
		}
	}

	if input.GetConfirm("Configure GitHub integration?", true) {
		fmt.Println()

		setGitHubConfig(gh)
	}

	return gh
}

func setGitHubConfig(gh *thirdParty.GitHub) {
	gh.AccessToken = strings.TrimSpace(input.GetInput(
		"Github Access Token:",
		"",
		gh.AccessToken,
		survey.Required,
	))
	gh.WebhooksSecret = strings.TrimSpace(input.GetInput(
		"Github Webhooks Secret:",
		"",
		gh.WebhooksSecret,
		survey.Required,
	))
	gh.Enabled = true
}
