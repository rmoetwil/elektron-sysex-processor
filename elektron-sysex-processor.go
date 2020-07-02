package main

import (
	"flag"
	"fmt"
	"github.com/rmoetwil/elektron-sysex-processor/sysex"
)

func main() {
	fileName := flag.String("sysex", "mm010308.syx", "The SysEx file to read")

	sysExData := sysex.ReadSysExFile(fileName)

	//For debugging dump part of the file to console
	sysExData.DumpSysExBytesToConsole(0, 1300)

	// Get some stats on the file
	messages := sysExData.GetAllMessages()

	fmt.Printf("Read %d messages \n", len(messages))
	fmt.Println("Processing finished")
}
