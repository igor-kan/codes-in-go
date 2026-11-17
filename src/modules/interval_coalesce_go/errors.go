package mod_interval_coalesce_go

import "errors"

var (
	ErrIntervalInvalid   = errors.New("interval: invalid argument")
	ErrIntervalBounds    = errors.New("interval: index out of bounds")
	ErrIntervalEmpty     = errors.New("interval: empty container")
	ErrIntervalInvariant = errors.New("interval: invariant violated")
)
