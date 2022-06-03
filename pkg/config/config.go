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
	// versionStable string = "stable"
	versionDev string = "dev"

	defaultAPIServerAuthServer    string = "https://mmesh.network"
	defaultAPIServerAuthServerDev string = "https://dev.mmesh.network"
	defaultAPIServerEndpoint      string = "mmesh.network:443"
	defaultAPIServerEndpointDev   string = "dev.mmesh.network:443"

	defaultServiceID string = "mmesh"
)

func Init() {
	hostID, err := os.Hostname()
	if err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	mmID := mmid.NewMMCLIID(viper.GetString("account.id"), hostID)
	viper.Set("mm.id", mmID.String())

	viper.Set("mm.app", version.CLI_NAME)

	viper.Set("host.id", hostID)

	SetDefaults()

	logging.LogLevel = xlog.GetLogLevel(viper.GetString("loglevel"))
	if logging.LogLevel == -1 {
		logging.LogLevel = xlog.INFO
	}

	logging.Interactive = true
}

func SetDefaults() {
	var isDev bool

	if viper.GetString("version.branch") == versionDev {
		isDev = true
	}

	if os.Getenv("MMESH_VERSION") == versionDev {
		viper.Set("version.branch", versionDev)
		isDev = true
	}

	if isDev {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServerDev)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpointDev)
	} else {
		viper.Set("apiserver.authServer", defaultAPIServerAuthServer)
		viper.Set("apiserver.endpoint", defaultAPIServerEndpoint)
	}

	viper.Set("serviceID", defaultServiceID)
}
