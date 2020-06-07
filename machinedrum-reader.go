package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readSysExFileAndDumpBytesToConsole(sysExFile *string) {
	sysex, err := ioutil.ReadFile(*sysExFile)
	check(err)

	fmt.Printf("Read %d bytes \n", len(sysex))

	//Hex dump
	//stdoutDumper := hex.Dumper(os.Stdout)
	//defer stdoutDumper.Close()
	//
	//stdoutDumper.Write(sysex)
}

func main() {
	sysExFile := flag.String("sysex", "md010308.syx", "The SysEx file to read")

	fmt.Println("Reading ", *sysExFile)

	readSysExFileAndDumpBytesToConsole(sysExFile)

	fmt.Println("Processing finished")

}
