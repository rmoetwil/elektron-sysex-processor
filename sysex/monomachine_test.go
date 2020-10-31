package sysex

import "testing"

func TestValidateMonoMachineSysEx_valid(t *testing.T) {
	if !ValidateMonoMachineSysEx([]byte{0xF0, 0x00, 0x20, 0x3c, 0x03, 0x00}) {
		t.Error("Not a monomachine id")
	}
}

func TestValidateMonoMachineSysEx_invalid(t *testing.T) {
	if ValidateMonoMachineSysEx([]byte{0xF0, 0x00, 0x20, 0x3c, 0x02, 0x00}) {
		t.Error("A monomachine id")
	}
}
