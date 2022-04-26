package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/update"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Update client and show version information",
	Long:  appHeader("Update client and show version information."),
}

// versionShowCmd represents the version/show command
var versionShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show " + version.CLI_NAME + " version information",
	Long:  appHeader("Show " + version.CLI_NAME + " version information"),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		header()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client Info: " + version.CLI_NAME + " " + version.GetVersion() + "\n")
	},
}

// versionUpdateCmd represents the version/update command
var versionUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update " + version.CLI_NAME + " to the latest version",
	Long:  appHeader("Update " + version.CLI_NAME + " to the latest version"),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		header()
	},
	Run: func(cmd *cobra.Command, args []string) {
		output.Header("Software Update")

		if err := update.Update(version.CLI_NAME); err != nil {
			status.Error(err, "Update failed")
		}

		msg.Ok("Latest version installed")
	},
}

func init() {
	versionCmd.AddCommand(versionShowCmd)
	versionCmd.AddCommand(versionUpdateCmd)
}
