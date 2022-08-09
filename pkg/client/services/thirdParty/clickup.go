package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) ClickUp(c *thirdParty.ClickUp) *thirdParty.ClickUp {
	fmt.Println()

	output.Header("PagerDuty Setup")

	if c == nil {
		c = &thirdParty.ClickUp{}
	}

	if c.Enabled {
		if input.GetConfirm("Delete ClickUp integration?", false) {
			c = &thirdParty.ClickUp{}
		}
	}

	if input.GetConfirm("Configure ClickUp integration?", true) {
		fmt.Println()

		setClickupConfig(c)
	}

	return c
}

func setClickupConfig(c *thirdParty.ClickUp) {
	c.ApiKey = strings.TrimSpace(input.GetInput(
		"ClickUp API Key:",
		"",
		c.ApiKey,
		survey.Required,
	))
	c.SettingsListURL = strings.TrimSpace(input.GetInput(
		"ClickUp Settings List URL:",
		"",
		c.SettingsListURL,
		survey.Required,
	))
	c.Enabled = true
}
