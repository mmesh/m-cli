package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) Scaleway(scw *thirdParty.Scaleway) *thirdParty.Scaleway {
	output.Header("Scaleway Cloud Setup")

	if scw == nil {
		scw = &thirdParty.Scaleway{}
	}

	if scw.Enabled {
		if input.GetConfirm("Delete Scaleway Cloud integration?", false) {
			scw = &thirdParty.Scaleway{}
		}
	}

	if input.GetConfirm("Configure Scaleway Cloud integration?", true) {
		fmt.Println()

		setScalewayConfig(scw)
	}

	return scw
}

func setScalewayConfig(scw *thirdParty.Scaleway) {
	scw.OrganizationID = strings.TrimSpace(input.GetInput(
		"Scaleway Cloud OrganizationID:",
		"",
		scw.OrganizationID,
		survey.Required,
	))
	scw.ProjectID = strings.TrimSpace(input.GetInput(
		"Scaleway Cloud ProjectID:",
		"",
		scw.ProjectID,
		survey.Required,
	))
	scw.AccessKey = strings.TrimSpace(input.GetInput(
		"Scaleway Cloud Access Key:",
		"",
		string(scw.AccessKey),
		survey.Required,
	))
	scw.SecretKey = strings.TrimSpace(input.GetInput(
		"Scaleway Cloud Secret Key:",
		"",
		string(scw.SecretKey),
		survey.Required,
	))
	scw.Enabled = true
}
