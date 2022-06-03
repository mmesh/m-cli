package alert

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) NewNote(a *events.Alert) {
	if a == nil {
		a = getAlert()
	}

	userEmail := viper.GetString("user.email")
	if len(userEmail) == 0 {
		userEmail = input.GetInput("User Email:", "", "", input.ValidEmail)
	}

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	acr := &events.AlertNewCommentRequest{
		AccountID: a.AccountID,
		AlertID:   a.AlertID,
		UserEmail: userEmail,
		Text:      input.GetMultiline("Note:", "", "", survey.Required),
	}

	s := output.Spinner()

	_, err := nxc.NewAlertComment(context.TODO(), acr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to add alert note")
	}

	s.Stop()

	fmt.Printf(`

Your input has been sent!

Thanks for using mmesh!

`)
}
