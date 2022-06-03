package auth

import (
	"fmt"
	"os"
	"time"

	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) LoginRequired(accountID string) bool {
	if len(accountID) == 0 {
		status.Error(fmt.Errorf("missing accountID"), "Invalid accountID")
	}

	apiKeyFile, err := auth.GetAPIKeyFile(accountID)
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
