package sysex

import "testing"

func TestValidateMachineDrumSysEx_valid(t *testing.T) {
	if !ValidateMachineDrumSysEx([]byte{0xF0, 0x00, 0x20, 0x3c, 0x02, 0x00}) {
		t.Error("Not a machine drum id")
	}
}

func TestValidateMachineDrumSysEx_invalid(t *testing.T) {
	if ValidateMachineDrumSysEx([]byte{0xF0, 0x00, 0x20, 0x3c, 0x03, 0x00}) {
		t.Error("A machine drum id")
	}
}
