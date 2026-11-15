package mod_vector_clock_causality

import "errors"

var (
	ErrVClockEmpty     = errors.New("vclock: empty or nil input")
	ErrVClockBounds    = errors.New("vclock: index or dimension out of bounds")
	ErrVClockInvalid   = errors.New("vclock: invalid configuration or state")
	ErrVClockInvariant = errors.New("vclock: structural invariant violated")
)
