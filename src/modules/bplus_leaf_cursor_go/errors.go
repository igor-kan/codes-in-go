package mod_bplus_leaf_cursor_go

import "errors"

var (
	ErrBPlusCursorInvalid   = errors.New("bplus_leaf_cursor_go: invalid argument")
	ErrBPlusCursorBounds    = errors.New("bplus_leaf_cursor_go: index out of bounds")
	ErrBPlusCursorEmpty     = errors.New("bplus_leaf_cursor_go: empty input")
	ErrBPlusCursorInvariant = errors.New("bplus_leaf_cursor_go: invariant violated")
)
