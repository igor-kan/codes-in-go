package mod_token_bucket_limiter_d56

import "errors"

var (
	ErrLimiterInvalid   = errors.New("token_bucket_limiter_d56: invalid argument")
	ErrLimiterBounds    = errors.New("token_bucket_limiter_d56: index out of bounds")
	ErrLimiterEmpty     = errors.New("token_bucket_limiter_d56: empty input")
	ErrLimiterInvariant = errors.New("token_bucket_limiter_d56: invariant violated")
)
