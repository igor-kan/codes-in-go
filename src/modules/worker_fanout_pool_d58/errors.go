package mod_worker_fanout_pool_d58

import "errors"

var (
	ErrPoolInvalid   = errors.New("worker_fanout_pool_d58: invalid argument")
	ErrPoolBounds    = errors.New("worker_fanout_pool_d58: index out of bounds")
	ErrPoolEmpty     = errors.New("worker_fanout_pool_d58: empty input")
	ErrPoolInvariant = errors.New("worker_fanout_pool_d58: invariant violated")
)
