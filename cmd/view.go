package cmd

import (
	"fmt"

	"github.com/rmoetwil/elektron-sysex-processor/internal/sysex"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a SysEx file",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := viper.GetString("SYSEX_FILE")
		sysExData := sysex.ReadSysExFile(fileName)

		//For debugging dump part of the file to console
		sysExData.DumpSysExBytesToConsole(0, 1300)

		// Get some stats on the file
		messages := sysExData.GetAllMessages()

		fmt.Printf("Read %d messages \n", len(messages))
		fmt.Println("Processing finished")

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
