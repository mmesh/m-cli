package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) Slack(slck *thirdParty.Slack) *thirdParty.Slack {
	output.Header("Slack Setup")

	if slck == nil {
		slck = &thirdParty.Slack{}
	}

	if slck.Enabled {
		if input.GetConfirm("Delete Slack integration?", false) {
			slck = &thirdParty.Slack{}
		}
	}

	if input.GetConfirm("Configure Slack integration?", true) {
		fmt.Println()

		setSlackConfig(slck)
	}

	return slck
}

func setSlackConfig(slck *thirdParty.Slack) {
	slck.AlertsWebhook = strings.TrimSpace(input.GetInput(
		"Alert Notifications Webhook:",
		"",
		slck.AlertsWebhook,
		survey.Required,
	))
	slck.OpsWebhook = strings.TrimSpace(input.GetInput(
		"General Notifications Webhook:",
		"",
		slck.OpsWebhook,
		survey.Required,
	))
	slck.Enabled = true
}
