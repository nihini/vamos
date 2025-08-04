package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var command2Sub2Cmd = &cobra.Command{
	Use:   "sub2",
	Short: "A sub-command of Command 2",
	Run: func(cmd *cobra.Command, args []string) {
		str := viper.GetString("myapp.sub2.str")
		num := viper.GetInt("myapp.sub2.num")
		fmt.Printf("Sub2 Values are %s:%d\n", str, num)
	},
}

func init() {
	command2Cmd.AddCommand(command2Sub2Cmd)

	command2Sub2Cmd.Flags().String("str", "No Sub Message", "Sub command message")
	command2Sub2Cmd.Flags().Int("num", 0, "Sub command number")
	viper.BindPFlag("myapp.sub2.str", command2Sub2Cmd.Flags().Lookup("str"))
	viper.BindPFlag("myapp.sub2.num", command2Sub2Cmd.Flags().Lookup("num"))
}
