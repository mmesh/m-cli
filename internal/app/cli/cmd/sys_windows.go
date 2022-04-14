//go:build windows
// +build windows

package cmd

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	// "github.com/mitchellh/go-homedir"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

const defaultConfigFile string = `.mmesh\mmeshctl.yml`

func consoleInit() error {
	if runtime.GOOS == "windows" {
		if err := colors.EnableWindowsVirtualTerminalProcessing(); err != nil {
			return err
		}
	}

	return nil
}

func defaultConfigFileHelp() string {
	return `configuration file (default: $HOME\.mmesh\mmeshctl.yml)`
}

func getConfigFile(cfgFile string) string {
	if len(cfgFile) == 0 {
		// Find home directory.
		user, err := user.Current()
		if err != nil {
			msg.Error(err)
			os.Exit(1)
		}

		// home, err := homedir.Dir()
		// if err != nil {
		// 	msg.Error(err)
		// 	os.Exit(1)
		// }

		// Search config in home directory with name ".mmesh" (without extension).
		// viper.AddConfigPath(home)
		// configName := strings.TrimSuffix(config.DefaultConfigFile, filepath.Ext(config.DefaultConfigFile))
		// viper.SetConfigName(configName)

		cfgFile = fmt.Sprintf(`%s\%s`, user.HomeDir, defaultConfigFile)
	}

	return cfgFile
}
