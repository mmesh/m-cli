package auth

import (
	"os"
	"path/filepath"
	"time"

	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) LoginRequired() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		status.Error(err, "Unable to get home directory")
	}

	mmeshDir := filepath.Join(homeDir, ".mmesh")

	if err := os.MkdirAll(mmeshDir, 0700); err != nil {
		status.Error(err, "Unable to create .mmesh config directory")
	}

	apikeyFile := filepath.Join(mmeshDir, "apikey")

	apikeyInfo, err := os.Stat(apikeyFile)
	if os.IsNotExist(err) {
		return true
	}

	if time.Since(apikeyInfo.ModTime()) > time.Duration(time.Hour) {
		return true
	}

	return false
}

/*
func (api *API) AutoLogin(req *auth.LoginRequest) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		status.Error(err, "Unable to get home directory")
	}

	mmeshDir := filepath.Join(homeDir, ".mmesh")

	if err := os.MkdirAll(mmeshDir, 0700); err != nil {
		status.Error(err, "Unable to create .mmesh config directory")
	}

	apikeyFile := filepath.Join(mmeshDir, "apikey")

	apikeyInfo, err := os.Stat(apikeyFile)
	if os.IsNotExist(err) {
		api.Login(req, true)
		return
	}

	if time.Since(apikeyInfo.ModTime()) > time.Duration(time.Hour) {
		// msg.Info("Refreshing auth session..")
		api.Login(req, true)
	}
}
*/
