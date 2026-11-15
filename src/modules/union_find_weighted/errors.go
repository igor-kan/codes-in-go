package mod_union_find_weighted

import "errors"

var (
	ErrWDSUEmpty     = errors.New("wdsu: empty or nil input")
	ErrWDSUBounds    = errors.New("wdsu: index or dimension out of bounds")
	ErrWDSUInvalid   = errors.New("wdsu: invalid configuration or state")
	ErrWDSUInvariant = errors.New("wdsu: structural invariant violated")
)
