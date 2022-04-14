package auth

import "mmesh.dev/m-api-go/grpc/resources/iam/auth"

type Interface interface {
	LoginRequired() bool
	// AutoLogin(req *auth.LoginRequest)
	Login(req *auth.LoginRequest, verbose bool)
	Logout()
	PasswordReset()
	ConfirmationMailResend()
	// Token()

	SSHAuth() bool
}
type API struct{}

func Resource() Interface {
	return &API{}
}
