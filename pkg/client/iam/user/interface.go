package user

import (
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/client/iam/user/output"
)

type Interface interface {
	List()
	Show()
	Create()
	Delete()
	SetPermissions()
	Enable()
	Disable()

	ShowLoggedUser(loginReq *auth.LoginRequest)
	SetEmail(loginReq *auth.LoginRequest)
	SetPassword(loginReq *auth.LoginRequest)
	SetTOTP(loginReq *auth.LoginRequest)
	GetSSHKey(loginReq *auth.LoginRequest) *iam.SSHKey
	SetSSHKeys(loginReq *auth.LoginRequest)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
