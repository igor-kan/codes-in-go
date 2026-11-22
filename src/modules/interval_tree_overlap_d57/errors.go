package mod_interval_tree_overlap_d57

import "errors"

var (
	ErrSweepInvalid   = errors.New("interval_tree_overlap_d57: invalid argument")
	ErrSweepBounds    = errors.New("interval_tree_overlap_d57: index out of bounds")
	ErrSweepEmpty     = errors.New("interval_tree_overlap_d57: empty input")
	ErrSweepInvariant = errors.New("interval_tree_overlap_d57: invariant violated")
)
