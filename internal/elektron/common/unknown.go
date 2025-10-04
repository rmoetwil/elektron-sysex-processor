package common

import "fmt"

type UnknownMessage struct {
	Raw []byte
}

func (u *UnknownMessage) Type() string        { return "Unknown" }
func (u *UnknownMessage) Decode([]byte) error { return nil }
func (u *UnknownMessage) String() string      { return fmt.Sprintf("Unknown message: % X", u.Raw) }
