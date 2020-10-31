package sysex

import (
	"bytes"
	"fmt"
)

type monomachineHandler struct {
}

func (m *monomachineHandler) modelName() string {
	return "Elektron MonoMachine"
}

func (m *monomachineHandler) parseMessage(message SysExMessage) {
	fmt.Println("Implement Monomachine message parsing")
}

func getMonoMachineId() []byte {
	return []byte{0x00, 0x20, 0x3c, 0x03, 0x00}
}

func ValidateMonoMachineSysEx(sysExData []byte) bool {
	if bytes.HasPrefix(sysExData[1:], getMonoMachineId()) {
		fmt.Println("Valid MonoMachine SysEx")
		return true
	} else {
		fmt.Println("Invalid MonoMachine SysEx")
		return false
	}
}
