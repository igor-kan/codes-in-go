package mod_min_heap_scheduler_d59

import "errors"

var (
	ErrSchedInvalid   = errors.New("min_heap_scheduler_d59: invalid argument")
	ErrSchedBounds    = errors.New("min_heap_scheduler_d59: index out of bounds")
	ErrSchedEmpty     = errors.New("min_heap_scheduler_d59: empty input")
	ErrSchedInvariant = errors.New("min_heap_scheduler_d59: invariant violated")
)
