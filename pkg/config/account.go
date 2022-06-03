package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
)

func DefaultAccount(cAuthServer, cEndpoint, accountID, userEmail string) error {
	if len(cAuthServer) == 0 {
		return fmt.Errorf("invalid controllerAuthServer")
	}

	if len(cEndpoint) == 0 {
		return fmt.Errorf("invalid controllerEndpoint")
	}

	if len(accountID) == 0 {
		return fmt.Errorf("invalid accountID")
	}

	if len(userEmail) == 0 {
		return fmt.Errorf("invalid userEmail")
	}

	if len(viper.GetString("account.id")) > 0 &&
		len(viper.GetString("user.email")) > 0 &&
		len(viper.GetString("controller.authServer")) > 0 &&
		len(viper.GetString("controller.endpoint")) > 0 {
		// account config already exists in mmeshctl.yml
		return nil
	}

	if input.GetConfirm("Want to configure this account as your default?", true) {
		writeCLIConfig(viper.ConfigFileUsed(), cAuthServer, cEndpoint, accountID, userEmail)
	}

	return nil
}

func writeCLIConfig(filename, cAuthServer, cEndpoint, accountID, userEmail string) {
	cfg := `# mmeshctl configuration

version:
  branch: ` + os.Getenv("MMESH_VERSION") + `

controller:
  authServer: ` + cAuthServer + `
  endpoint: ` + cEndpoint + `

account:
  id: ` + accountID + `

user:
  email: ` + userEmail + `

agent:
  management:
    auth:
      psk:
      securityToken:

`

	dir := fmt.Sprintf("%s/%s", filepath.Dir(filename), accountID)
	if err := os.MkdirAll(dir, 0700); err != nil {
		status.Error(err, "Unable to create config directory")
	}

	if err := ioutil.WriteFile(filename, []byte(cfg), 0600); err != nil {
		status.Error(err, "Unable to write configuration")
	}
}
