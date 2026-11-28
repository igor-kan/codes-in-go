package mod_brent_cycle_detect_go_d61

import "errors"

var (
	ErrBrentCycleInvalid   = errors.New("brent_cycle_detect_go_d61: invalid argument")
	ErrBrentCycleBounds    = errors.New("brent_cycle_detect_go_d61: index out of bounds")
	ErrBrentCycleEmpty     = errors.New("brent_cycle_detect_go_d61: empty input")
	ErrBrentCycleInvariant = errors.New("brent_cycle_detect_go_d61: invariant violated")
)
