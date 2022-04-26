package grpc

import (
	"os"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/grpc/client"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

type apiClientParameters struct {
	authKey    *auth_pb.AuthKey
	authSecret string
	grpcServer string
}

func getManagerAPIClientParams(reqAuth bool) *apiClientParameters {
	p := &apiClientParameters{}

	if reqAuth {
		authKey, err := auth.GetAuthKey()
		if err != nil {
			msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
			os.Exit(1)
		}
		p.authKey = authKey
	} else {
		p.authKey = auth.GetNoAuthKey()
	}

	// p.authSecret = viper.GetString("controller.authSecret")
	p.grpcServer = viper.GetString("apiserver.endpoint")

	return p
}
func getControllerAPIClientParams(reqAuth bool) *apiClientParameters {
	p := &apiClientParameters{}

	if reqAuth {
		authKey, err := auth.GetAuthKey()
		if err != nil {
			msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
			os.Exit(1)
		}
		p.authKey = authKey
	} else {
		p.authKey = auth.GetNoAuthKey()
	}

	// p.authSecret = viper.GetString("controller.authSecret")

	p.grpcServer = viper.GetString("controller.endpoint")
	if len(p.grpcServer) == 0 {
		msg.Error("Missing configuration. Check config file.")
		os.Exit(1)
	}

	return p
}

func GetBillingAPIClient(reqAuth bool) (rpc.BillingAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewBillingAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetServicesAPIClient(reqAuth bool) (rpc.ServicesAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewServicesAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetManagerProviderAPIClient(reqAuth bool) (rpc.ProviderAPIClient, *grpc.ClientConn) {
	p := getManagerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewProviderAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetControllerProviderAPIClient(reqAuth bool) (rpc.ProviderAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams(reqAuth)

	nxc, conn, err := client.NewProviderAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetCoreAPIClient() (rpc.CoreAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams(true)

	nxc, conn, err := client.NewCoreAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		// msg.Errorf("Unable to connect to gRPC server. Check configuration and connectivity: %v", err)
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}

func GetNetworkAPIClient() (rpc.NetworkAPIClient, *grpc.ClientConn) {
	p := getControllerAPIClientParams(true)

	nxc, conn, err := client.NewNetworkAPIClient(p.grpcServer, p.authKey, p.authSecret)
	if err != nil {
		msg.Errorf("Unable to connect to gRPC server: %v", errors.Cause(err))
		os.Exit(1)
	}

	return nxc, conn
}
