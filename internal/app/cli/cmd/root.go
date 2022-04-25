package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/internal/app/cli/setup"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/config"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/update"
	"mmesh.dev/m-lib/pkg/utils"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
)

var cfgFile string
var newVersionAvailable bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   version.CLI_NAME,
	Short: version.CLI_NAME + " is a CLI to control the " + version.PLATFORM_NAME,
	Long: version.CLI_NAME + ` is a CLI to control the ` +
		version.PLATFORM_NAME + `.

Find support and more information:

  Project Website:     ` + version.MMESH_URL + `
  Documentation:       ` + version.MMESH_DOC_URL + `
  Join Us on Discord:  ` + version.DISCORD_URL,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// silent auto-login
	autoLogin()

	if err := rootCmd.Execute(); err != nil {
		msg.Error(errors.Cause(err))
		os.Exit(1)
	}

	if newVersionAvailable {
		fmt.Printf("%s\n", colors.DarkBlue("_"))
		cmd := colors.DarkWhite(fmt.Sprintf("%s version update", version.CLI_NAME))
		q := colors.DarkBlue("'")
		msg := fmt.Sprintf("%s %s%s%s", colors.Black("New version available, please update with"), q, cmd, q)
		fmt.Printf("%s %s\n\n", colors.Cyan("🢂"), msg)
	}
}

func init() {
	if err := consoleInit(); err != nil {
		msg.Error(err)
		os.Exit(1)
	}

	fmt.Println(colors.Black(version.CLI_NAME + " " + version.GetVersion()))

	// in order to work with 'sudo' needs to be called before initConfig()
	swUpdate()

	initConfig()

	go checkVersionUpdate()

	output.AppHeader(vars.AccountID, false)
	// fmt.Printf("%s %s\n\n", version.CLI_NAME, version.GetVersion())

	cobra.EnableCommandSorting = false

	// cobra.OnInitialize(initConfig, adminOpts)
	cobra.OnInitialize()

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// rootCmd.CompletionOptions.DisableDescriptions = true

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", defaultConfigFileHelp())
	// rootCmd.PersistentFlags().BoolP("insecure", "", false, "insecure mode (TLS disabled)")
	// rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format (table, json)")
	rootCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	flag.StringVar(&vars.AccountID, "a", "", "account identifier")
	flag.StringVar(&cfgFile, "config", "", defaultConfigFileHelp())
	flag.Parse()

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	cfgFile = getConfigFile(cfgFile)

	viper.SetConfigFile(cfgFile)

	viper.SetEnvPrefix("mm") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv() // read in environment variables that match

	if !utils.FileExists(cfgFile) {
		config.SetDefaults()
		if err := setup.Setup(cfgFile); err != nil {
			msg.Errorf("Unable to create config file: %v", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		msg.Errorf("Unable to read config file: %v", err)
		os.Exit(1)
	}
	// msg.Debugf("Using configuration file: %v", viper.ConfigFileUsed())

	config.Init()
}

func checkVersionUpdate() {
	newVersionAvailable, _ = update.IsBinaryOutdated(version.CLI_NAME)
}

func swUpdate() {
	if len(os.Args) == 3 {
		if os.Args[1] == "version" && os.Args[2] == "update" {
			output.SectionHeader("Software Update")

			s := output.Spinner()

			if err := update.Update(version.CLI_NAME); err != nil {
				s.Stop()
				status.Error(err, "Update failed")
			}

			s.Stop()

			msg.Ok("Latest version installed")
			os.Exit(0)
		}
	}
}

func autoLogin() {
	if len(os.Args) > 1 {
		if os.Args[1] == "auth" {
			return
		}
	}

	if client.Auth().LoginRequired() {
		client.Auth().Login(login.NewRequest(), true)
	}
	// client.Auth().AutoLogin(login.NewRequest())
}
