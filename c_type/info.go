package c_type

import "C"
import (
	"encoding/hex"
	"fmt"
)

type Einfo [INFO_WIDTH]byte

func (b Einfo) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b[:])
	return result, nil
}

func (b *Einfo) UnmarshalText(input []byte) error {
	if len(input) < 2 {
		return fmt.Errorf("hex string length must > 2 : current is %v", len(input))
	}
	raw := input[2:]
	if len(raw) == 0 {
		return nil
	}
	dec := Einfo{}
	if len(raw)/2 != len(dec[:]) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(dec[:])*2, "Einfo")
	}
	if _, err := hex.Decode(dec[:], raw); err != nil {
		return err
	} else {
		*b = dec
	}
	return nil
}
