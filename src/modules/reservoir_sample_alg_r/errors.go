package mod_reservoir_sample_alg_r

import "errors"

var (
	ErrReservoirGoInvalid   = errors.New("reservoir_sample_alg_r: invalid argument")
	ErrReservoirGoBounds    = errors.New("reservoir_sample_alg_r: index out of bounds")
	ErrReservoirGoEmpty     = errors.New("reservoir_sample_alg_r: empty input")
	ErrReservoirGoInvariant = errors.New("reservoir_sample_alg_r: invariant violated")
)
