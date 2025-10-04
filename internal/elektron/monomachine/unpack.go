package monomachine

import (
	"fmt"

	"github.com/rmoetwil/elektron-sysex-processor/internal/elektron/common"
)

// Should contain Monomachine-specific unpacking routines—things that are unique to how Monomachine encodes or structures its payloads beyond the standard MIDI 7-bit packing.

// Examples (to be verified with specs

// Field level decoders
func decodeTempo(data []byte) float64 {
	// Machinedrum tempo might be stored as two bytes, e.g. BPM * 10
	return float64(data[0]) + float64(data[1])/10
}

func decodeFlags(flagByte byte) (clockIn, ctrlIn bool) {
	clockIn = flagByte&0x01 != 0
	ctrlIn = flagByte&0x02 != 0
	return
}

// Bitfield Extraction (Machinedrum uses packed flags and mode bits:)

func extractModeBits(b byte) (mode int, enabled bool) {
	mode = int((b >> 1) & 0x03)
	enabled = b&0x01 != 0
	return
}

// Checksum validation (If Machinedrum messages include checksums:)

func validateChecksum(data []byte, expected byte) bool {
	sum := byte(0)
	for _, b := range data {
		sum += b
	}
	return sum == expected
}

// Message-specific Unpackers
func unpackPatternData(raw []byte) ([]byte, error) {
	// Use common.unpack7Bit first, then apply Machinedrum-specific slicing
	unpacked := common.Unpack7Bit(raw)
	//TODO How to get the expected length
	expectedLength := len(unpacked)
	if len(unpacked) < expectedLength {
		return nil, fmt.Errorf("pattern too short")
	}
	return unpacked, nil
}
