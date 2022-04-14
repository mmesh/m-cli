package itsm

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *issueAPI) NewComment(i *itsm.Issue) {
	if i == nil {
		i = getIssue(false)
	}

	userEmail := viper.GetString("user.email")
	if len(userEmail) == 0 {
		userEmail = input.GetInput("User Email:", "", "", input.ValidEmail)
	}

	userNickname := strings.Split(userEmail, "@")[0]

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	icr := &itsm.IssueNewCommentRequest{
		AccountID:    i.AccountID,
		IssueID:      i.IssueID,
		ServiceID:    i.ServiceID,
		ProviderID:   i.ProviderID,
		UserEmail:    userEmail,
		UserNickname: userNickname,
		Comment:      input.GetMultiline("New comment:", "", "", survey.Required),
	}

	s := output.Spinner()

	_, err := nxc.NewIssueComment(context.TODO(), icr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to add issue comment")
	}

	s.Stop()

	fmt.Printf(`

Your input has been sent!

Thanks for using mmesh!

`)
}
