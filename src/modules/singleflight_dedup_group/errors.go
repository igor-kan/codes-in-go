package mod_singleflight_dedup_group

import "errors"

var (
	ErrFlightEmpty     = errors.New("flight: empty or nil input")
	ErrFlightBounds    = errors.New("flight: index or dimension out of bounds")
	ErrFlightInvalid   = errors.New("flight: invalid configuration or state")
	ErrFlightInvariant = errors.New("flight: structural invariant violated")
)
