package mod_consistent_hash_vnodes_d56

import "errors"

var (
	ErrRingInvalid   = errors.New("consistent_hash_vnodes_d56: invalid argument")
	ErrRingBounds    = errors.New("consistent_hash_vnodes_d56: index out of bounds")
	ErrRingEmpty     = errors.New("consistent_hash_vnodes_d56: empty input")
	ErrRingInvariant = errors.New("consistent_hash_vnodes_d56: invariant violated")
)
