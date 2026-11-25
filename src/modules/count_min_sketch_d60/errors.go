package mod_count_min_sketch_d60

import "errors"

var (
	ErrCMSketchInvalid   = errors.New("count_min_sketch_d60: invalid argument")
	ErrCMSketchBounds    = errors.New("count_min_sketch_d60: index out of bounds")
	ErrCMSketchEmpty     = errors.New("count_min_sketch_d60: empty input")
	ErrCMSketchInvariant = errors.New("count_min_sketch_d60: invariant violated")
)
