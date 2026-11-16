package mod_rw_fair_rwlock_queue

import "errors"

var (
	ErrFairRWEmpty     = errors.New("fairrw: empty or nil input")
	ErrFairRWBounds    = errors.New("fairrw: index or dimension out of bounds")
	ErrFairRWInvalid   = errors.New("fairrw: invalid configuration or state")
	ErrFairRWInvariant = errors.New("fairrw: structural invariant violated")
)
