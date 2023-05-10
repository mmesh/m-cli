package login

import (
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func NewRequest() *auth.LoginRequest {
	userToken := viper.GetString("token")

	if len(userToken) == 0 {
		msg.Error("Authorization token not found")
		os.Exit(1)
	}

	return &auth.LoginRequest{
		UserToken: userToken,
	}
}
