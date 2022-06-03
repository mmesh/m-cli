package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) PagerDuty(pd *thirdParty.PagerDuty) *thirdParty.PagerDuty {
	fmt.Println()

	output.Header("PagerDuty Setup")

	if pd == nil {
		pd = &thirdParty.PagerDuty{}
	}

	if pd.Enabled {
		if input.GetConfirm("Delete PagerDuty integration?", false) {
			pd = &thirdParty.PagerDuty{}
		}
	}

	if input.GetConfirm("Configure PagerDuty integration?", true) {
		fmt.Println()

		setPagerDutyConfig(pd)
	}

	return pd
}

func setPagerDutyConfig(pd *thirdParty.PagerDuty) {
	pd.IntegrationKey = strings.TrimSpace(input.GetInput(
		"PagerDuty Integration Key:",
		"",
		pd.IntegrationKey,
		survey.Required,
	))
	pd.Enabled = true
}
