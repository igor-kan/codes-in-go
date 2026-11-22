package mod_count_min_sketch_d57

import "errors"

var (
	ErrCmsInvalid   = errors.New("count_min_sketch_d57: invalid argument")
	ErrCmsBounds    = errors.New("count_min_sketch_d57: index out of bounds")
	ErrCmsEmpty     = errors.New("count_min_sketch_d57: empty input")
	ErrCmsInvariant = errors.New("count_min_sketch_d57: invariant violated")
)
