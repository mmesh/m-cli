package auth

import (
	"os"

	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Logout() {
	apiKeyFile, err := auth.GetAPIKeyFile()
	if err != nil {
		status.Error(err, "Unable to find API key")
	}

	if err := os.RemoveAll(apiKeyFile); err != nil {
		status.Error(err, "Unable to logout")
	}

	output.Logout()
}
