package mod_reservoir_heap_topk_d57

import "errors"

var (
	ErrTopKInvalid   = errors.New("reservoir_heap_topk_d57: invalid argument")
	ErrTopKBounds    = errors.New("reservoir_heap_topk_d57: index out of bounds")
	ErrTopKEmpty     = errors.New("reservoir_heap_topk_d57: empty input")
	ErrTopKInvariant = errors.New("reservoir_heap_topk_d57: invariant violated")
)
