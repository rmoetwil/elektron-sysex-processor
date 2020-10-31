package sysex

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	SysExStart byte = 0xF0
	SysExEnd   byte = 0xF7
)

type sysExMessageHandler interface {
	modelName() string
	parseMessage(SysExMessage)
}

type SysExModel struct {
	model   []byte
	handler sysExMessageHandler
}

var sysExModels = []SysExModel{
	{getMachineDrumId(), &machinedrumHandler{}},
	{getMonoMachineId(), &monomachineHandler{}},
}

type SysExData struct {
	bytes []byte
}

type SysExMessage struct {
	bytes []byte
}

func (message SysExMessage) IsValid() bool {
	return message.bytes[0] == SysExStart && message.bytes[len(message.bytes)-1] == SysExEnd
}

// TODO Move this to the specific handler
func (message SysExMessage) Type() byte {
	sysExModel := message.Model()
	if sysExModel != nil {
		return message.bytes[1+len(sysExModel.model)]
	} else {
		// TODO error?!?!?
		return 0
	}
}

func (message SysExMessage) Model() *SysExModel {
	for _, model := range sysExModels {
		if bytes.HasPrefix(message.bytes[1:], model.model) {
			return &model
		}
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadSysExFile(sysExFile *string) SysExData {
	fmt.Println("Reading ", *sysExFile)

	sysExData, err := ioutil.ReadFile(*sysExFile)
	check(err)

	fmt.Printf("Read %d bytes \n", len(sysExData))

	return SysExData{sysExData}
}

func (data SysExData) DumpSysExBytesToConsole(start, length int) {
	//Hex dump
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	stdoutDumper.Write(data.bytes[start : start+length])
}

func (data SysExData) GetAllMessages() []SysExMessage {
	messages := make([]SysExMessage, 0)

	messageStart := 0
	sysExData := data.bytes
	for i := 0; i < len(sysExData)-1; i++ {
		if sysExData[i] == SysExEnd {
			// Create message
			message := SysExMessage{sysExData[messageStart : i+1]}

			sysExModel := message.Model()
			fmt.Printf("%s message %x valid %t \n", sysExModel.handler.modelName(), message.Type(), message.IsValid())

			sysExModel.handler.parseMessage(message)
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
