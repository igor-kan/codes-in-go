package mod_b_tree_order_m

import "errors"

var (
	ErrBTreeEmpty     = errors.New("btree: empty or nil input")
	ErrBTreeBounds    = errors.New("btree: index or dimension out of bounds")
	ErrBTreeInvalid   = errors.New("btree: invalid configuration or state")
	ErrBTreeInvariant = errors.New("btree: structural invariant violated")
)
