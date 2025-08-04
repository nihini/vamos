package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd is the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-cobra-viper",
	Short: "CLI for managing services",
	Long:  "A reusable, stable CLI template using Cobra and Viper",
}

// Execute adds all child commands to the root and sets flags
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-cobra-viper.yaml)")
	rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Environment variable support
	viper.SetEnvPrefix("CLI")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("cli-cobra-viper")
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
