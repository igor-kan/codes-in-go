package mod_fenwick_range_point_go

import "errors"

var (
	ErrBitGoInvalid   = errors.New("fenwick_range_point_go: invalid argument")
	ErrBitGoBounds    = errors.New("fenwick_range_point_go: index out of bounds")
	ErrBitGoEmpty     = errors.New("fenwick_range_point_go: empty input")
	ErrBitGoInvariant = errors.New("fenwick_range_point_go: invariant violated")
)
