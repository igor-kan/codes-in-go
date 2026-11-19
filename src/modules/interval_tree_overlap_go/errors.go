package mod_interval_tree_overlap_go

import "errors"

var (
	ErrIvTreeGoInvalid   = errors.New("interval_tree_overlap_go: invalid argument")
	ErrIvTreeGoBounds    = errors.New("interval_tree_overlap_go: index out of bounds")
	ErrIvTreeGoEmpty     = errors.New("interval_tree_overlap_go: empty input")
	ErrIvTreeGoInvariant = errors.New("interval_tree_overlap_go: invariant violated")
)
