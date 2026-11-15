package mod_swiss_probe_hashmap

import "errors"

var (
	ErrSwissEmpty     = errors.New("swiss: empty or nil input")
	ErrSwissBounds    = errors.New("swiss: index or dimension out of bounds")
	ErrSwissInvalid   = errors.New("swiss: invalid configuration or state")
	ErrSwissInvariant = errors.New("swiss: structural invariant violated")
)
