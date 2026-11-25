package mod_suffix_array_kasai_d60

import "errors"

var (
	ErrSuffixIndexInvalid   = errors.New("suffix_array_kasai_d60: invalid argument")
	ErrSuffixIndexBounds    = errors.New("suffix_array_kasai_d60: index out of bounds")
	ErrSuffixIndexEmpty     = errors.New("suffix_array_kasai_d60: empty input")
	ErrSuffixIndexInvariant = errors.New("suffix_array_kasai_d60: invariant violated")
)
