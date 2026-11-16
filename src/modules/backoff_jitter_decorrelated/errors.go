package mod_backoff_jitter_decorrelated

import "errors"

var (
	ErrBackoffEmpty     = errors.New("backoff: empty or nil input")
	ErrBackoffBounds    = errors.New("backoff: index or dimension out of bounds")
	ErrBackoffInvalid   = errors.New("backoff: invalid configuration or state")
	ErrBackoffInvariant = errors.New("backoff: structural invariant violated")
)
