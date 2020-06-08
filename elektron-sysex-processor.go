package main

import (
	"flag"
	"fmt"
	"github.com/rmoetwil/elektron-sysex-processor/sysex"
)

func main() {
	sysExFile := flag.String("sysex", "md010308.syx", "The SysEx file to read")
	sysExType := flag.String("type", "machinedrum", "The SysEx file type")

	sysExData := sysex.ReadSysExFileAndDumpBytesToConsole(sysExFile)

	sysex.ValidateGenericSysEx(sysExData)

	//TODO refactor this implementation
	if *sysExType == "machinedrum" {
		sysex.ValidateMachineDrumSysEx(sysExData)
	}

	if *sysExType == "monomachine" {
		sysex.ValidateMonoMachineSysEx(sysExData)
	}
	fmt.Println("Processing finished")
}
