package mod_lfu_cache_o1_buckets

import "errors"

var (
	ErrLFUEmpty     = errors.New("lfu: empty or nil input")
	ErrLFUBounds    = errors.New("lfu: index or dimension out of bounds")
	ErrLFUInvalid   = errors.New("lfu: invalid configuration or state")
	ErrLFUInvariant = errors.New("lfu: structural invariant violated")
)
