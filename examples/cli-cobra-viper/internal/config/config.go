package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// AppConfig holds all configuration values
type AppConfig struct {
	Verbose bool
	MyApp   struct {
		Str     string
		Num     int
		Message string
		Sub2    struct {
			Str string
			Num int
		}
	}
	API struct {
		Endpoint string
		Token    string
	}
}

// Config is the global config instance
var Config AppConfig

// InitConfig reads config file, env vars, flags, and unmarshals into Config
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting home dir: %v\n", err)
		}
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("cli-cobra-viper")
	}

	viper.SetConfigType("yaml")

	// Set defaults for all config fields
	viper.SetDefault("verbose", false)
	viper.SetDefault("myapp.str", "A string")
	viper.SetDefault("myapp.num", 42)
	viper.SetDefault("myapp.message", "No Message")
	viper.SetDefault("myapp.sub2.str", "No Sub Message")
	viper.SetDefault("myapp.sub2.num", 0)
	viper.SetDefault("api.endpoint", "https://api.example.com")
	viper.SetDefault("api.token", "")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Unmarshal Viper config into AppConfig struct
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding config: %v\n", err)
	}
}
