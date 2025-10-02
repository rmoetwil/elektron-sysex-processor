package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "sysex-processor",
	Short: "SysEx CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand like 'view'")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	viper.AutomaticEnv()
}
