package auth

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/errors"
)

func setApiKey(authKey *auth.AuthKey) error {
	jsonData, err := json.Marshal(authKey)
	if err != nil {
		return errors.Wrapf(err, "[%v] function json.Marshal()", errors.Trace())
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.UserHomeDir()", errors.Trace())
	}

	mmeshDir := filepath.Join(homeDir, ".mmesh")

	if err := os.MkdirAll(mmeshDir, 0700); err != nil {
		return errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
	}

	mmeshApiKeyPath := filepath.Join(mmeshDir, "apikey")

	if err := ioutil.WriteFile(mmeshApiKeyPath, jsonData, 0600); err != nil {
		return errors.Wrapf(err, "[%v] function ioutil.WriteFile()", errors.Trace())
	}

	return nil
}
