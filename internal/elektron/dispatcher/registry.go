package dispatcher

import (
	"bytes"

	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/common"
	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/machinedrum"
	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/monomachine"
)

func getMachineDrumId() []byte {
	return []byte{0xF0, 0x00, 0x20, 0x3C, 0x02, 0x00}
}

func getMonoMachineId() []byte {
	return []byte{0xF0, 0x00, 0x20, 0x3C, 0x03, 0x00}
}

func isMonomachine(msg []byte) bool {
	return len(msg) > 6 && bytes.HasPrefix(msg, getMonoMachineId())
}

func isMachinedrum(msg []byte) bool {
	return len(msg) > 6 && bytes.HasPrefix(msg, getMachineDrumId())
}

func RouteMessage(data []byte) common.SysExMessage {
	if isMonomachine(data) {
		return monomachine.NewMessage(data)
	} else if isMachinedrum(data) {
		return machinedrum.NewMessage(data)
	}
	return &common.UnknownMessage{Raw: data}
}
