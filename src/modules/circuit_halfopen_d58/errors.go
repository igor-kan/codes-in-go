package mod_circuit_halfopen_d58

import "errors"

var (
	ErrBreakerInvalid   = errors.New("circuit_halfopen_d58: invalid argument")
	ErrBreakerBounds    = errors.New("circuit_halfopen_d58: index out of bounds")
	ErrBreakerEmpty     = errors.New("circuit_halfopen_d58: empty input")
	ErrBreakerInvariant = errors.New("circuit_halfopen_d58: invariant violated")
)
