package mod_backoff_jitter_decorrelated

import "time"

// DelayPolicy generates sequential backoff intervals across retry attempts.
type DelayPolicy interface {
	Next(attempt int) time.Duration
	Reset()
}
