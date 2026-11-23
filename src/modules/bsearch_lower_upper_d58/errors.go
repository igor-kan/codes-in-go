package mod_bsearch_lower_upper_d58

import "errors"

var (
	ErrBSearchInvalid   = errors.New("bsearch_lower_upper_d58: invalid argument")
	ErrBSearchBounds    = errors.New("bsearch_lower_upper_d58: index out of bounds")
	ErrBSearchEmpty     = errors.New("bsearch_lower_upper_d58: empty input")
	ErrBSearchInvariant = errors.New("bsearch_lower_upper_d58: invariant violated")
)
