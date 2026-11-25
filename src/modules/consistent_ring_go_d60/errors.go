package mod_consistent_ring_go_d60

import "errors"

var (
	ErrHashRingInvalid   = errors.New("consistent_ring_go_d60: invalid argument")
	ErrHashRingBounds    = errors.New("consistent_ring_go_d60: index out of bounds")
	ErrHashRingEmpty     = errors.New("consistent_ring_go_d60: empty input")
	ErrHashRingInvariant = errors.New("consistent_ring_go_d60: invariant violated")
)
