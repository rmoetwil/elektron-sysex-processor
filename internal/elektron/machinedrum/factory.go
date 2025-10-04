package machinedrum

import (
	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/common"
)

const (
	MsgGlobal  = 0x50
	MsgPattern = 0x52
	MsgKit     = 0x54
)

func NewMessage(data []byte) Message {
	if len(data) < 7 {
		return &common.UnknownMessage{Raw: data}
	}

	msgID := data[6] // Machinedrum message ID byte

	switch msgID {
	case MsgGlobal:
		return &GlobalSettings{}
	//case MsgPattern:
	//	return &PatternDump{}
	//case MsgKit:
	//	return &KitDump{}
	default:
		return &common.UnknownMessage{Raw: data}
	}
}
