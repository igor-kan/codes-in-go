package mod_lfu_o1_cache_d56

import "errors"

var (
	ErrLfuInvalid   = errors.New("lfu_o1_cache_d56: invalid argument")
	ErrLfuBounds    = errors.New("lfu_o1_cache_d56: index out of bounds")
	ErrLfuEmpty     = errors.New("lfu_o1_cache_d56: empty input")
	ErrLfuInvariant = errors.New("lfu_o1_cache_d56: invariant violated")
)
