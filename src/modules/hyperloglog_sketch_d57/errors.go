package mod_hyperloglog_sketch_d57

import "errors"

var (
	ErrHllInvalid   = errors.New("hyperloglog_sketch_d57: invalid argument")
	ErrHllBounds    = errors.New("hyperloglog_sketch_d57: index out of bounds")
	ErrHllEmpty     = errors.New("hyperloglog_sketch_d57: empty input")
	ErrHllInvariant = errors.New("hyperloglog_sketch_d57: invariant violated")
)
