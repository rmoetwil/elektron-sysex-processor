package machinedrum

import (
	"fmt"

	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/common"
)

type GlobalSettings struct {
	MidiChannel byte
	ClockIn     bool
	// ...
}

func (g *GlobalSettings) Type() string {
	return "GlobalSettings"
}

func (g *GlobalSettings) Decode(data []byte) error {
	unpacked := common.Unpack7Bit(data)
	g.MidiChannel = unpacked[0]
	g.ClockIn = unpacked[5]&0x01 != 0
	return nil
}

func (g *GlobalSettings) String() string {
	return fmt.Sprintf("GlobalSettings: MidiChannel=%d, ClockIn=%v", g.MidiChannel, g.ClockIn)
}
