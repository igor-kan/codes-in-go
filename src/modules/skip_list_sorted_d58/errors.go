package mod_skip_list_sorted_d58

import "errors"

var (
	ErrSkipChunkInvalid   = errors.New("skip_list_sorted_d58: invalid argument")
	ErrSkipChunkBounds    = errors.New("skip_list_sorted_d58: index out of bounds")
	ErrSkipChunkEmpty     = errors.New("skip_list_sorted_d58: empty input")
	ErrSkipChunkInvariant = errors.New("skip_list_sorted_d58: invariant violated")
)
