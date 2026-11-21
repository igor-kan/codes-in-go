package mod_bloom_double_hash_d56

import "errors"

var (
	ErrBloomInvalid   = errors.New("bloom_double_hash_d56: invalid argument")
	ErrBloomBounds    = errors.New("bloom_double_hash_d56: index out of bounds")
	ErrBloomEmpty     = errors.New("bloom_double_hash_d56: empty input")
	ErrBloomInvariant = errors.New("bloom_double_hash_d56: invariant violated")
)
