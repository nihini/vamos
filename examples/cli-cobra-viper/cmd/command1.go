package cmd

import (
	"fmt"

	"github.com/nihini/cli-cobra-viper/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cmd1 represents the base command for Command 1
var command1 = &cobra.Command{
	Use:   "cmd1",
	Short: "Test Command 1",
}

// Sub Command of Command 1
// This command will display a message passed as an argument and also read a message from the config
// It demonstrates how to use viper to bind flags and read configuration values
var command1_Sub1 = &cobra.Command{
	Use:   "sub [message]",
	Short: "Sub Command of Command 1 - display a message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := config.Config.MyApp.Message
		fmt.Printf("Sub Command 1 - Message from argument: %s - Message from config: %s\n", args[0], message)
	},
}

func init() {
	rootCmd.AddCommand(command1)
	command1.AddCommand(command1_Sub1)

	command1_Sub1.Flags().String("message", "No Message", "message to display")
	viper.BindPFlag("myapp.message", command1_Sub1.Flags().Lookup("message"))
}
