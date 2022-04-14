package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
)

func DefaultAccount(accountID, userEmail string) error {
	if len(accountID) == 0 || accountID == "admin" || accountID == "master" {
		return nil
	}

	if len(viper.GetString("account.id")) > 0 && len(viper.GetString("user.email")) > 0 {
		return nil
	}

	if input.GetConfirm("Want to configure this account as your default?", true) {
		viper.Set("account.id", accountID)
		viper.Set("user.email", userEmail)

		WriteCLIConfig(viper.ConfigFileUsed())
	}

	return nil
}

func WriteCLIConfig(filename string) {
	cfg := `# mmeshctl configuration

version:
  branch: ` + os.Getenv("MMESH_VERSION") + `

controller:
  authServer: ` + viper.GetString("controller.authServer") + `
  endpoint: ` + viper.GetString("controller.endpoint") + `

account:
  id: ` + viper.GetString("account.id") + `

user:
  email: ` + viper.GetString("user.email") + `

agent:
  management:
    auth:
      psk:
      securityToken:

`

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0700); err != nil {
		status.Error(err, "Unable to create config directory")
	}

	if err := ioutil.WriteFile(filename, []byte(cfg), 0600); err != nil {
		status.Error(err, "Unable to write configuration")
	}
}
