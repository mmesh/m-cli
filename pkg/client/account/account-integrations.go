package account

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	thirdParty_pb "mmesh.dev/m-api-go/grpc/resources/services/thirdParty"
	"mmesh.dev/m-cli/pkg/client/services/thirdParty"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Settings() {
	a := FetchAccount()

	Output().Configuration(a)

	if !input.GetConfirm("Edit account configuration now?", false) {
		fmt.Println()
		return
	}

	nxc, grpcConn := grpc.GetControllerProviderAPIClient(true)
	defer grpcConn.Close()

	fmt.Println()

	a.Description = input.GetInput(
		"Description:",
		"",
		a.Description,
		survey.Required,
	)

	if a.Integrations == nil {
		a.Integrations = &thirdParty_pb.Integrations{
			Github:    &thirdParty_pb.GitHub{},
			Pagerduty: &thirdParty_pb.PagerDuty{},
			Slack:     &thirdParty_pb.Slack{},
			// Crisp:     &thirdParty_pb.Crisp{},

			DigitalOcean: &thirdParty_pb.DigitalOcean{},
			GCP:          &thirdParty_pb.GCP{},
			Scaleway:     &thirdParty_pb.Scaleway{},
			// AWS:          &thirdParty_pb.AWS{},

		}
	}

	a.Integrations.Github = thirdParty.Setup().GitHub(a.Integrations.Github)
	a.Integrations.Pagerduty = thirdParty.Setup().PagerDuty(a.Integrations.Pagerduty)
	a.Integrations.Slack = thirdParty.Setup().Slack(a.Integrations.Slack)
	// a.Integrations.Crisp = thirdParty.Setup().Crisp(a.Integrations.Crisp)
	a.Integrations.DigitalOcean = thirdParty.Setup().DigitalOcean(a.Integrations.DigitalOcean)
	a.Integrations.GCP = thirdParty.Setup().GCP(a.Integrations.GCP)
	a.Integrations.Scaleway = thirdParty.Setup().Scaleway(a.Integrations.Scaleway)
	// a.Integrations.AWS = thirdParty.Setup().AWS(a.Integrations.AWS)

	s := output.Spinner()

	a, err := nxc.SetAccountIntegrations(context.TODO(), a)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set account")
	}

	s.Stop()

	Output().Configuration(a)
}
