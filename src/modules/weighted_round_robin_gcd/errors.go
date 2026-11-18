package mod_weighted_round_robin_gcd

import "errors"

var (
	ErrSmoothWrrInvalid   = errors.New("weighted_round_robin_gcd: invalid argument")
	ErrSmoothWrrBounds    = errors.New("weighted_round_robin_gcd: index out of bounds")
	ErrSmoothWrrEmpty     = errors.New("weighted_round_robin_gcd: empty input")
	ErrSmoothWrrInvariant = errors.New("weighted_round_robin_gcd: invariant violated")
)
