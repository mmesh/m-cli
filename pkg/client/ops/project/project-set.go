package project

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/ops/project"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	p := GetProject(true)
	if p != nil { // editing existing resource
		output.Choice("Edit Project")
	} else { // <new> resource
		output.Choice("New Project")

		p = &project.Project{
			AccountID: a.AccountID,
		}

		if err := survey.AskOne(
			&survey.Input{Message: "Project ID:"},
			&p.ProjectID,
			survey.WithValidator(validProjectID),
			survey.WithIcons(input.SurveySetIcons),
		); err != nil {
			status.Error(err, "Unable to get response")
		}
	}

	p.Description = input.GetInput("Description:", "", p.Description, survey.Required)

	p.ReviewRequired = input.GetConfirm("Enable workflow-required reviews?", p.ReviewRequired)

	p.ApprovalRequired = input.GetConfirm("Enable workflow-required approvals?", p.ApprovalRequired)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	p, err := nxc.SetProject(context.TODO(), p)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set project")
	}

	s.Stop()

	Output().Show(p)
}
