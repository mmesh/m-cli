package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) Crisp(crisp *thirdParty.Crisp) *thirdParty.Crisp {
	output.Header("Crisp Setup")

	if crisp == nil {
		crisp = &thirdParty.Crisp{}
	}

	if crisp.Enabled {
		if input.GetConfirm("Delete Crisp integration?", false) {
			crisp = &thirdParty.Crisp{}
		}
	}

	if input.GetConfirm("Configure Crisp integration?", true) {
		fmt.Println()

		setCrispConfig(crisp)
	}

	return crisp
}

func setCrispConfig(crisp *thirdParty.Crisp) {
	crisp.WebsiteID = strings.TrimSpace(input.GetInput(
		"Crisp WebsiteID:",
		"",
		crisp.WebsiteID,
		survey.Required,
	))
	crisp.Enabled = true
}
