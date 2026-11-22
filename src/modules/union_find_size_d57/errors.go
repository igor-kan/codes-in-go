package mod_union_find_size_d57

import "errors"

var (
	ErrDsuSizeInvalid   = errors.New("union_find_size_d57: invalid argument")
	ErrDsuSizeBounds    = errors.New("union_find_size_d57: index out of bounds")
	ErrDsuSizeEmpty     = errors.New("union_find_size_d57: empty input")
	ErrDsuSizeInvariant = errors.New("union_find_size_d57: invariant violated")
)
