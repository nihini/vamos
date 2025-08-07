package cmd

import (
	"strings"

	"github.com/nihini/cli-cobra-viper/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// ...AppConfig and Config are now in internal/config.go...

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
	cobra.OnInitialize(func() { config.InitConfig(cfgFile) })

	// Global persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-cobra-viper.yaml)")
	rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Environment variable support
	viper.SetEnvPrefix("CLI")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// ...initConfig now lives in internal/config.go...
