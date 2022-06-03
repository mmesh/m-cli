package auth

import "mmesh.dev/m-api-go/grpc/resources/iam/auth"

type Interface interface {
	LoginRequired(accountID string) bool
	// AutoLogin(req *auth.LoginRequest)
	Login(req *auth.LoginRequest, verbose bool)
	Logout(accountID string)
	PasswordReset()
	ConfirmationMailResend()
	// Token()

	SSHAuth(accountID string) bool
}
type API struct{}

func Resource() Interface {
	return &API{}
}
