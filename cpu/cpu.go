package cpu

import "fmt"

// Status holds data for each status flag
// in the 6502 status register.
type Status struct {
	C bool // carry
	Z bool // zero result
	I bool // interrupt disable
	D bool // decimal mode
	B bool // break command
	V bool // overflow
	N bool // zero result
}

// String implements Stringer
func (sr *Status) String() string {
	convert := func(bit bool) string {
		if bit {
			return "1"
		}
		return "0"
	}

	return fmt.Sprintf("%s%s0%s%s%s%s%s",
		convert(sr.N),
		convert(sr.V),
		convert(sr.B),
		convert(sr.D),
		convert(sr.I),
		convert(sr.Z),
		convert(sr.C),
	)
}

// Clear clears all bits of the status register,
// i.e. will set all flags to '0'.
func (sr *Status) Clear() {
	sr.C = false
	sr.Z = false
	sr.I = false
	sr.D = false
	sr.B = false
	sr.V = false
	sr.N = false
}

// Registers holds data for each register
// used by the 6502,
type Registers struct {
	Status      *Status
	Accumulator byte
}
