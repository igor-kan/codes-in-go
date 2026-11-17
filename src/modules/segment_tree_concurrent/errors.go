package mod_segment_tree_concurrent

import "errors"

var (
	ErrSegTreeInvalid   = errors.New("segtree: invalid argument")
	ErrSegTreeBounds    = errors.New("segtree: index out of bounds")
	ErrSegTreeEmpty     = errors.New("segtree: empty container")
	ErrSegTreeInvariant = errors.New("segtree: invariant violated")
)
