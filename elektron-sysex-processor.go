package main

import (
	"flag"
	"fmt"
	"github.com/rmoetwil/elektron-sysex-processor/sysex"
)

func main() {
	sysExFile := flag.String("sysex", "mm010308.syx", "The SysEx file to read")
	sysExType := flag.String("type", "machinedrum", "The SysEx file type")

	sysExData := sysex.ReadSysExFileAndDumpBytesToConsole(sysExFile)

	//TODO refactor this implementation
	if *sysExType == "machinedrum" {
		sysex.ValidateMachineDrumSysEx(sysExData)
	}

	if *sysExType == "monomachine" {
		sysex.ValidateMonoMachineSysEx(sysExData)
	}

	messages := sysex.GetAllMessages(sysExData)

	for _, message := range messages {
		fmt.Printf("%x\n", message.Message)

	}

	fmt.Println("Processing finished")
}
