package mod_sliding_percentile_p2_go

import "errors"

var (
	ErrWindowQuantInvalid   = errors.New("sliding_percentile_p2_go: invalid argument")
	ErrWindowQuantBounds    = errors.New("sliding_percentile_p2_go: index out of bounds")
	ErrWindowQuantEmpty     = errors.New("sliding_percentile_p2_go: empty input")
	ErrWindowQuantInvariant = errors.New("sliding_percentile_p2_go: invariant violated")
)
