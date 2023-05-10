package auth

import (
	"os"
	"time"

	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) LoginRequired() bool {
	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		status.Error(err, "Unable to find API key")
	}

	apiKeyInfo, err := os.Stat(apiKeyFile)
	if os.IsNotExist(err) {
		return true
	}

	if time.Since(apiKeyInfo.ModTime()) > time.Hour {
		return true
	}

	return false
}
