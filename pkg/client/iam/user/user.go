package user

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

type loggedUser struct {
	accountID string
	email     string
	isAdmin   bool
}

func GetUser(edit bool) *iam.User {
	ul := users()

	if len(ul) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var userOptID string
	usersOpts := make([]string, 0)
	users := make(map[string]*iam.User)

	for _, u := range ul {
		userOptID = u.Email
		usersOpts = append(usersOpts, userOptID)
		users[userOptID] = u
	}

	sort.Strings(usersOpts)

	if edit {
		usersOpts = append(usersOpts, input.NewResource)
	}

	userOptID = input.GetSelect("User:", "", usersOpts, survey.Required)

	if userOptID == input.NewResource {
		return nil
	}

	vars.UserEmail = users[userOptID].Email

	return users[userOptID]
}

func users() map[string]*iam.User {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &iam.ListUsersRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	users := make(map[string]*iam.User)

	for {
		ul, err := nxc.ListUsers(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list users")
		}

		for _, u := range ul.Users {
			users[u.Email] = u
		}

		if len(ul.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = ul.Meta.NextPageToken
		} else {
			break
		}
	}

	return users
}

func getLoggedUser() *loggedUser {
	isAdmin := viper.GetBool("logged.isAdmin")
	realm := viper.GetString("logged.realm")
	userEmail := viper.GetString("logged.email")

	if len(realm) == 0 || len(userEmail) == 0 {
		status.Error(fmt.Errorf("user not logged in"), "Unable to get user data")
	}

	return &loggedUser{
		accountID: realm,
		email:     userEmail,
		isAdmin:   isAdmin,
	}
}
