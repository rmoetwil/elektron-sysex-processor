package sysex

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	SysExStart byte = 0xF0
	SysExEnd   byte = 0xF7
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ValidateGenericSysEx(sysExData []byte) {
	if sysExData[0] != SysExStart || sysExData[len(sysExData)-1] != SysExEnd {
		fmt.Println("Sysex not valid")
	} else {
		fmt.Println("Sysex valid")
	}
}

func ReadSysExFileAndDumpBytesToConsole(sysExFile *string) []byte {
	fmt.Println("Reading ", *sysExFile)

	sysExData, err := ioutil.ReadFile(*sysExFile)
	check(err)

	fmt.Printf("Read %d bytes \n", len(sysExData))

	//fmt.Printf("%T\n", sysex)

	//Hex dump
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	stdoutDumper.Write(sysExData[:256])

	return sysExData
}
