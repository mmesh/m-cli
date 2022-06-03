package itsm

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

type IssueInterface interface {
	New()
	List()
	Show() *itsm.Issue
	Delete()
	Close()
	NewComment(i *itsm.Issue)
}
type issueAPI struct{}

func (api *API) Issue() IssueInterface {
	return &issueAPI{}
}

/*
func submitIssue(issueType itsm.IssueType, issueSubtype itsm.IssueSubtype) {
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

	switch issueType {
	case itsm.IssueType_FEEDBACK:
		issueSubtype = getFeedbackType()
	case itsm.IssueType_INCIDENT:
		issueSubtype = getIncidentType()
	}

	i.IssueType = issueType
	i.IssueSubtype = issueSubtype
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
*/

func getIssue(hideClosed bool) *itsm.Issue {
	issues := issues()

	if len(issues) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	issuesOpts := make([]string, 0)
	issueMap := make(map[string]*itsm.Issue)

	for _, i := range issues {
		tm := output.Datetime(i.LastModified)
		var status string
		if i.Status == itsm.IssueStatus_ISSUE_CLOSED {
			if hideClosed {
				continue
			}
			status = "CLOSED"
		} else {
			status = "OPEN"
		}
		summary := output.Fit(i.Summary, 37)
		issueOptID := fmt.Sprintf("%s │ %s [%s]", tm, summary, status)
		issuesOpts = append(issuesOpts, issueOptID)
		issueMap[issueOptID] = i
	}

	if len(issuesOpts) == 0 {
		msg.Info("No open tickets found")
		os.Exit(1)
	}

	issueID := input.GetSelect("Ticket:", "", issuesOpts, survey.Required)

	vars.IssueID = issueMap[issueID].IssueID

	return issueMap[issueID]
}

func issues() map[string]*itsm.Issue {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &itsm.ListIssuesRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	issues := make(map[string]*itsm.Issue)

	for {
		il, err := nxc.ListIssues(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list ITSM issues")
		}

		for _, i := range il.Issues {
			issues[i.IssueID] = i
		}

		if len(il.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = il.Meta.NextPageToken
		} else {
			break
		}
	}

	return issues
}
