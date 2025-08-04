package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var command2Cmd = &cobra.Command{
	Use:   "cmd2",
	Short: "Test Command 2",
	Run: func(cmd *cobra.Command, args []string) {
		str := viper.GetString("myapp.str")
		num := viper.GetInt("myapp.num")
		fmt.Printf("Values are [%s] and [%d]\n", str, num)
	},
}

func init() {
	rootCmd.AddCommand(command2Cmd)

	command2Cmd.Flags().String("str", "A string", "A string value")
	command2Cmd.Flags().Int("num", 42, "A number value")
	viper.BindPFlag("myapp.str", command2Cmd.Flags().Lookup("str"))
	viper.BindPFlag("myapp.num", command2Cmd.Flags().Lookup("num"))
}
