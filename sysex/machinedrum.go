package sysex

import (
	"bytes"
	"fmt"
)

func getMachineDrumId() []byte {
	return []byte{0x00, 0x20, 0x3c, 0x02, 0x00}
}

func ValidateMachineDrumSysEx(sysExData []byte) bool {
	if bytes.HasPrefix(sysExData[1:], getMachineDrumId()) {
		fmt.Println("Valid MachineDrum SysEx")
		return true
	} else {
		fmt.Println("Invalid MachineDrum SysEx")
		return false
	}
}
