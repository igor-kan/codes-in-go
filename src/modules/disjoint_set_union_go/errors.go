package mod_disjoint_set_union_go

import "errors"

var (
	ErrDsuGoInvalid   = errors.New("disjoint_set_union_go: invalid argument")
	ErrDsuGoBounds    = errors.New("disjoint_set_union_go: index out of bounds")
	ErrDsuGoEmpty     = errors.New("disjoint_set_union_go: empty input")
	ErrDsuGoInvariant = errors.New("disjoint_set_union_go: invariant violated")
)
