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

type SysExMessage struct {
	Header  []byte
	Message byte
	Data    []byte
}

func check(e error) {
	if e != nil {
		panic(e)
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

	stdoutDumper.Write(sysExData[:1300])

	return sysExData
}

func GetAllMessages(sysExData []byte) []SysExMessage {
	messages := make([]SysExMessage, 0)

	//Start reading messages.
	//Format is:
	// - F0 byte
	// - header 5 bytes
	// - message type 1 byte
	// - data
	// - F7 byte
	//TODO Elektron's MonoMachine and MachineDrum headers are indeed 5 bytes long. Others might not.
	headerLength := 5
	messageStart := 0
	for i := 0; i < len(sysExData)-1; i++ {
		if sysExData[i] == SysExEnd {
			// Create message
			message := SysExMessage{sysExData[messageStart+1 : messageStart+1+headerLength],
				sysExData[messageStart+1+headerLength],
				sysExData[messageStart+1+headerLength : i-1]}

			messages = append(messages, message)
			// Point to next message
			messageStart = i + 1

			if sysExData[messageStart] != SysExStart {
				//TODO Do something here as this is a signal the parsing is off.
				fmt.Println("Not at message start")
			}
		}
	}

	return messages
}
