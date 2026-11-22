package mod_sliding_log_limiter_d57

import "errors"

var (
	ErrSLogInvalid   = errors.New("sliding_log_limiter_d57: invalid argument")
	ErrSLogBounds    = errors.New("sliding_log_limiter_d57: index out of bounds")
	ErrSLogEmpty     = errors.New("sliding_log_limiter_d57: empty input")
	ErrSLogInvariant = errors.New("sliding_log_limiter_d57: invariant violated")
)
