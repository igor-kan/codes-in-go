package mod_crdt_lww_element_set

import "errors"

var (
	ErrLWWEmpty     = errors.New("lww: empty or nil input")
	ErrLWWBounds    = errors.New("lww: index or dimension out of bounds")
	ErrLWWInvalid   = errors.New("lww: invalid configuration or state")
	ErrLWWInvariant = errors.New("lww: structural invariant violated")
)
