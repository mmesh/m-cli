package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/utils"
)

// GetNoAuthKey gets void authKey
func GetNoAuthKey() *auth.AuthKey {
	return &auth.AuthKey{
		Key: "no-auth",
	}
}

// GetAuthKey gets the authorization bearer string key
func GetAuthKey() (*auth.AuthKey, error) {
	var authKey auth.AuthKey

	accountID := viper.GetString("account.id")
	if len(accountID) == 0 {
		return nil, fmt.Errorf("missing account.id in configuration file")
	}

	apiKeyFile, err := GetAPIKeyFile(accountID)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getAPIKeyFile()", errors.Trace())
	}

	jsonBlob, err := utils.ReadJsonFile(apiKeyFile)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function utils.ReadJsonFile()", errors.Trace())
	}

	if err := json.Unmarshal(jsonBlob, &authKey); err != nil {
		return nil, errors.Wrapf(err, "[%v] function json.Unmarshal(jsonBlob, &apiKey)", errors.Trace())
	}

	return &authKey, nil
}

func GetAPIKeyFile(accountID string) (string, error) {
	if len(accountID) == 0 {
		return "", fmt.Errorf("missing accountID")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function os.UserHomeDir()", errors.Trace())
	}

	mmeshAccountDir := filepath.Join(homeDir, ".mmesh", accountID)

	if err := os.MkdirAll(mmeshAccountDir, 0700); err != nil {
		return "", errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
	}

	return filepath.Join(mmeshAccountDir, "apikey"), nil
}

/*
// GetAuthToken removes user credentials
func GetAuthToken() error {
	var authToken auth.AuthToken
	var adminToken auth.AdminToken

	authKey, err := GetAuthKey()
	if err != nil {
		return errors.Wrapf(err, "[%v] function GetAuthBearer()", errors.Trace())
	}
	jsonData, err := base64.URLEncoding.DecodeString(authKey.Key)
	if err != nil {
		return errors.Wrapf(err, "[%v] function base64.URLEncoding.DecodeString(authKey.Key)", errors.Trace())
	}
	if err := json.Unmarshal(jsonData, &authToken); err != nil {
		return errors.Wrapf(err, "[%v] function json.Unmarshal(jsonData, &authToken)", errors.Trace())
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Target: ")
	target, _ := reader.ReadString('\n')

	adminToken.Email = authToken.Subkey1
	adminToken.Realm = authToken.Subkey2
	adminToken.Target = strings.TrimSpace(target)
	adminToken.Timestamp = time.Now().Unix()
	adminToken.ExpTime = adminToken.Timestamp + 1800

	jsonData, err = json.Marshal(adminToken)
	if err != nil {
		return errors.Wrapf(err, "[%v] function json.Marshal(adminToken)", errors.Trace())
	}

	token := base64.URLEncoding.EncodeToString(jsonData)

	fmt.Printf("\n%s: %s\n", colors.Cyan("Authorization Token"), colors.Black(token))

	return nil
}
*/
