package login

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/input"
)

func NewRequest() *auth.LoginRequest {
	var auto bool

	req := &auth.LoginRequest{}

	req.Realm = viper.GetString("account.id")
	req.Email = viper.GetString("user.email")

	if client.Auth().SSHAuth() && len(req.Realm) > 0 && len(req.Email) > 0 {
		auto = true
	}

	if !auto {
		req.Realm = input.GetInput("Account:", "", viper.GetString("account.id"), survey.Required)
		req.Email = input.GetInput("Email:", "", viper.GetString("user.email"), input.ValidEmail)
	}

	if client.Auth().SSHAuth() {
		req.AuthMethod = auth.AuthMethod_SSH_KEY
	} else {
		req.AuthMethod = auth.AuthMethod_PASSWORD
		req.Password = input.GetPassword("Password:", "")
	}

	return req
}
