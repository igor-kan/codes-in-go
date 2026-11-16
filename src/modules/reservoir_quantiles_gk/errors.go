package mod_reservoir_quantiles_gk

import "errors"

var (
	ErrQuantileEmpty     = errors.New("quantile: empty or nil input")
	ErrQuantileBounds    = errors.New("quantile: index or dimension out of bounds")
	ErrQuantileInvalid   = errors.New("quantile: invalid configuration or state")
	ErrQuantileInvariant = errors.New("quantile: structural invariant violated")
)
