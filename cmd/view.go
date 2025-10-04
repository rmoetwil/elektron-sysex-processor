package cmd

import (
	"fmt"

	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/common"
	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/dispatcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a SysEx file",
	Run: func(cmd *cobra.Command, args []string) {
		debug := viper.GetBool("DEBUG")
		fileName := viper.GetString("SYSEX_FILE")
		data := common.ReadSysExFile(fileName)

		if debug {
			//For debugging dump part of the file to console
			common.DumpSysExBytesToConsole(data, 0, 1300)
		}

		// Get some stats on the file
		messages := common.ExtractSysExMessages(data)

		if debug {
			for _, message := range messages {
				sysExMessage := dispatcher.RouteMessage(message)
				//fmt.Printf("%s:\n", sysExMessage.Type())
				fmt.Printf("%s:\n", sysExMessage.String())
			}
		}
		fmt.Println("Processing finished")

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
