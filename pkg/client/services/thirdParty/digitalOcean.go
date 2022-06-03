package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) DigitalOcean(dgo *thirdParty.DigitalOcean) *thirdParty.DigitalOcean {
	fmt.Println()

	output.Header("DigitalOcean Setup")

	if dgo == nil {
		dgo = &thirdParty.DigitalOcean{}
	}

	if dgo.Enabled {
		if input.GetConfirm("Delete DigitalOcean integration?", false) {
			dgo = &thirdParty.DigitalOcean{}
		}
	}

	if input.GetConfirm("Configure DigitalOcean integration?", true) {
		fmt.Println()

		setDigitalOceanConfig(dgo)
	}

	return dgo
}

func setDigitalOceanConfig(dgo *thirdParty.DigitalOcean) {
	dgo.Token = strings.TrimSpace(input.GetInput(
		"DigitalOcean Token:",
		"",
		dgo.Token,
		survey.Required,
	))
	// dgo.ProjectName = strings.TrimSpace(input.GetInput(
	// 	"DigitalOcean Project Name:",
	// 	"Projects are resource groups to organize your cloud infrastructure items",
	// 	dgo.ProjectName,
	// 	survey.Required,
	// ))
	dgo.Enabled = true
}
