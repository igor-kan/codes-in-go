package mod_consistent_bounded_load

import "errors"

var (
	ErrBoundedRingInvalid   = errors.New("boundedring: invalid argument")
	ErrBoundedRingBounds    = errors.New("boundedring: index out of bounds")
	ErrBoundedRingEmpty     = errors.New("boundedring: empty container")
	ErrBoundedRingInvariant = errors.New("boundedring: invariant violated")
)
