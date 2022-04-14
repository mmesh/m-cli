package itsm

import (
	"context"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

const (
	issueAssistanceInfo          string = "information request"
	issueAssistanceSupport       string = "technical support request"
	issueFeedbackGeneralComments string = "feedback / general comments"
	issueFeedbackFeatureRequest  string = "feature request"
	issueIncidentBilling         string = "billing issue"
	issueIncidentPlatform        string = "software / platform issue"
)

type issueType struct {
	issueType    itsm.IssueType
	issueSubtype itsm.IssueSubtype
}

func (api *issueAPI) New() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	i := &itsm.Issue{
		AccountID: a.AccountID,
	}

	i.OwnerUserEmail = viper.GetString("user.email")
	if len(i.OwnerUserEmail) == 0 {
		i.OwnerUserEmail = input.GetInput("User Email:", "", "", input.ValidEmail)
	}

	it := getIssueType()

	i.IssueType = it.issueType
	i.IssueSubtype = it.issueSubtype
	i.Summary = input.GetInput("Summary:", "", "", survey.Required)
	i.Description = input.GetInput("Description:", "", "", survey.Required)

	s := output.Spinner()

	i, err := nxc.NewIssue(context.TODO(), i)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to submit issue")
	}

	s.Stop()

	Output().Show(i)
}

func getIssueType() *issueType {
	issueTypes := map[string]*issueType{
		issueAssistanceInfo: {
			issueType:    itsm.IssueType_ASSISTANCE,
			issueSubtype: itsm.IssueSubtype_ASSISTANCE_INFO,
		},
		issueAssistanceSupport: {
			issueType:    itsm.IssueType_ASSISTANCE,
			issueSubtype: itsm.IssueSubtype_ASSISTANCE_SUPPORT,
		},
		issueFeedbackGeneralComments: {
			issueType:    itsm.IssueType_FEEDBACK,
			issueSubtype: itsm.IssueSubtype_FEEDBACK_GENERAL_COMMENTS,
		},
		issueFeedbackFeatureRequest: {
			issueType:    itsm.IssueType_FEEDBACK,
			issueSubtype: itsm.IssueSubtype_FEEDBACK_FEATURE_REQUEST,
		},
		issueIncidentBilling: {
			issueType:    itsm.IssueType_INCIDENT,
			issueSubtype: itsm.IssueSubtype_INCIDENT_BILLING,
		},
		issueIncidentPlatform: {
			issueType:    itsm.IssueType_INCIDENT,
			issueSubtype: itsm.IssueSubtype_INCIDENT_PLATFORM,
		},
	}

	iTypeOpts := make([]string, 0)
	for iType := range issueTypes {
		iOpt := strings.Title(iType)
		iTypeOpts = append(iTypeOpts, iOpt)
	}

	iType := input.GetSelect("Select category:", "", iTypeOpts, survey.Required)

	return issueTypes[strings.ToLower(iType)]
}
