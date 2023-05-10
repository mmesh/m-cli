package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/utils"
)

type Credentials struct {
	AccountID    string `json:"accountID,omitempty"`
	FederationID string `json:"federationID,omitempty"`
	Key          string `json:"key,omitempty"`
}

// GetNoAuthKey gets void authKey
func GetNoAuthKey() *auth.AuthKey {
	return &auth.AuthKey{
		Key: "no-auth",
	}
}

// GetAuthKey gets the authorization bearer string key
func GetAuthKey() (*auth.AuthKey, error) {
	cred, err := getCredentials()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	return &auth.AuthKey{Key: cred.Key}, nil
}

func GetAccountID() (string, error) {
	if len(vars.AccountID) > 0 {
		return vars.AccountID, nil
	}

	cred, err := getCredentials()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	if len(cred.AccountID) == 0 {
		return "", fmt.Errorf("missing accountID")
	}

	return cred.AccountID, nil
}

func getFederationID() (string, error) {
	if len(vars.FederationID) > 0 {
		return vars.FederationID, nil
	}

	cred, err := getCredentials()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	if len(cred.FederationID) == 0 {
		return "", fmt.Errorf("missing federationID")
	}

	return cred.FederationID, nil
}

func GetControllerEndpoint() (string, error) {
	federationID, err := getFederationID()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getFederationID()", errors.Trace())
	}

	apiserver := viper.GetString("apiserver.endpoint")
	if len(apiserver) == 0 {
		return "", fmt.Errorf("invalid apiserver endpoint")
	}

	return fmt.Sprintf("%s.%s", federationID, apiserver), nil
}

func GetAPIKeyFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function os.UserHomeDir()", errors.Trace())
	}

	mmeshDir := filepath.Join(homeDir, ".mmesh")

	if err := os.MkdirAll(mmeshDir, 0700); err != nil {
		return "", errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
	}

	return filepath.Join(mmeshDir, "apikey"), nil
}

func getCredentials() (*Credentials, error) {
	apiKeyFile, err := GetAPIKeyFile()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getAPIKeyFile()", errors.Trace())
	}

	jsonBlob, err := utils.ReadJsonFile(apiKeyFile)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function utils.ReadJsonFile()", errors.Trace())
	}

	var cred Credentials

	if err := json.Unmarshal(jsonBlob, &cred); err != nil {
		return nil, errors.Wrapf(err, "[%v] function json.Unmarshal()", errors.Trace())
	}

	return &cred, nil
}
