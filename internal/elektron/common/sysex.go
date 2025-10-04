package common

import (
	"encoding/hex"
	"fmt"
	"os"
)

const (
	SysExStart byte = 0xF0
	SysExEnd   byte = 0xF7
)

type SysExMessage interface {
	Type() string
	Decode([]byte) error
	String() string
}

func DumpSysExBytesToConsole(data []byte, start, length int) {
	//Hex dump
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	stdoutDumper.Write(data[start : start+length])
}

func ExtractSysExMessages(data []byte) [][]byte {
	var messages [][]byte
	i := 0
	for i < len(data) {
		if data[i] == SysExStart {
			start := i
			for i < len(data) && data[i] != SysExEnd {
				i++
			}
			if i < len(data) && data[i] == SysExEnd {
				end := i
				messages = append(messages, data[start:end+1])
			}
		}
		i++
	}
	return messages
}

func ReadSysExFile(fileName string) []byte {
	fmt.Println("Reading ", fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d bytes \n", len(data))

	return data
}
