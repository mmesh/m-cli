package config

import (
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

const (
	versionBranchStable string = "stable"
	versionBranchDev    string = "dev"

	defaultAPIServerAuthServer    string = "https://mmesh.network"
	defaultAPIServerAuthServerDev string = "https://dev.mmesh.network"
	defaultAPIServerEndpoint      string = "mmesh.network:443"
	defaultAPIServerEndpointDev   string = "dev.mmesh.network:443"

	defaultControllerAuthServer    string = "https://eu-central-01.mmesh.network"
	defaultControllerAuthServerDev string = "https://eu-central-01.dev.mmesh.network"
	defaultControllerEndpoint      string = "eu-central-01.mmesh.network:443"
	defaultControllerEndpointDev   string = "eu-central-01.dev.mmesh.network:443"

	defaultServiceID string = "mmesh"
)

func Init() {
	hostID, err := os.Hostname()
	if err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	//viper.Set("clientID", clientID)

	mmID := mmid.NewMMCLIID(hostID)

	viper.Set("mm.id", mmID.String())

	viper.Set("mm.app", version.CLI_NAME)

	viper.Set("host.id", hostID)

	SetDefaults()

	// logLevel
	logging.LogLevel = xlog.GetLogLevel(viper.GetString("loglevel"))
	if logging.LogLevel == -1 {
		logging.LogLevel = xlog.INFO
	}

	logging.Interactive = true
}

func SetDefaults() {
	var isDev bool

	if viper.GetString("version.branch") == versionBranchDev ||
		os.Getenv("MMESH_VERSION") == versionBranchDev {
		isDev = true
	}

	if isDev {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServerDev)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpointDev)
	} else {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServer)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpoint)
	}

	controllerAuthServer := viper.GetString("controller.authServer")
	if len(controllerAuthServer) == 0 {
		if isDev {
			viper.Set("controller.authServer", defaultControllerAuthServerDev)
		} else {
			viper.Set("controller.authServer", defaultControllerAuthServer)
		}
	}

	controllerEndpoint := viper.GetString("controller.endpoint")
	if len(controllerEndpoint) == 0 {
		if isDev {
			viper.Set("controller.endpoint", defaultControllerEndpointDev)
		} else {
			viper.Set("controller.endpoint", defaultControllerEndpoint)
		}
	}

	viper.Set("serviceID", defaultServiceID)
}
