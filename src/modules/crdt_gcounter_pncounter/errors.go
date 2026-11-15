package mod_crdt_gcounter_pncounter

import "errors"

var (
	ErrCRDTEmpty     = errors.New("crdt: empty or nil input")
	ErrCRDTBounds    = errors.New("crdt: index or dimension out of bounds")
	ErrCRDTInvalid   = errors.New("crdt: invalid configuration or state")
	ErrCRDTInvariant = errors.New("crdt: structural invariant violated")
)
