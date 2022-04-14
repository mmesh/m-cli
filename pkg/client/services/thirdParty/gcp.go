package thirdParty

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
)

func (a *API) GCP(gcp *thirdParty.GCP) *thirdParty.GCP {
	output.Header("Google Cloud Setup")

	if gcp == nil {
		gcp = &thirdParty.GCP{}
	}

	if gcp.Enabled {
		if input.GetConfirm("Delete Google Cloud integration?", false) {
			gcp = &thirdParty.GCP{}
		}
	}

	if input.GetConfirm("Configure Google Cloud integration?", true) {
		fmt.Println()

		setGCPConfig(gcp)
	}

	return gcp
}

func setGCPConfig(gcp *thirdParty.GCP) {
	gcp.ProjectID = strings.TrimSpace(input.GetInput(
		"Google Cloud Platform ProjectID:",
		"",
		gcp.ProjectID,
		survey.Required,
	))
	gcp.JSONKey = []byte(strings.TrimSpace(input.GetMultiline(
		"Google Cloud Platform JSON Key file:",
		"",
		string(gcp.JSONKey),
		survey.Required,
	)))
	gcp.Enabled = true
}
