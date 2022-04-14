package auth

import (
	"os"
	"path/filepath"

	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Logout() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		status.Error(err, "Unable to logout")
	}

	apiKeyFile := filepath.Join(homeDir, ".mmesh", "apikey")

	if err := os.RemoveAll(apiKeyFile); err != nil {
		status.Error(err, "Unable to logout")
	}

	output.Logout()
}
