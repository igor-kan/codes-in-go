package mod_count_min_heavy_go

import "errors"

var (
	ErrCountMinInvalid   = errors.New("countmin: invalid argument")
	ErrCountMinBounds    = errors.New("countmin: index out of bounds")
	ErrCountMinEmpty     = errors.New("countmin: empty container")
	ErrCountMinInvariant = errors.New("countmin: invariant violated")
)
