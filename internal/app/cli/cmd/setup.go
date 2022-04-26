package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mmesh.dev/m-cli/internal/app/cli/setup"
	"mmesh.dev/m-cli/pkg/vars"
)

// setupCmd represents the networkRoutes command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "First-time setup",
	Long:  appHeader(`First-time setup.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		header()
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
		cfgFile = getConfigFile(cfgFile)

		viper.SetConfigFile(cfgFile)

		viper.SetEnvPrefix("mm") // will be uppercased automatically
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.AutomaticEnv() // read in environment variables that match

		setup.Configure(cfgFile)
	},
}

func init() {
	setupCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
