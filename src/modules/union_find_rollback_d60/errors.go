package mod_union_find_rollback_d60

import "errors"

var (
	ErrRollbackDsuInvalid   = errors.New("union_find_rollback_d60: invalid argument")
	ErrRollbackDsuBounds    = errors.New("union_find_rollback_d60: index out of bounds")
	ErrRollbackDsuEmpty     = errors.New("union_find_rollback_d60: empty input")
	ErrRollbackDsuInvariant = errors.New("union_find_rollback_d60: invariant violated")
)
