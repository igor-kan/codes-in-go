package mod_backoff_jitter_decorrelated

import "time"

// JitterMode selects either FullJitter or Decorrelated backoff calculation.
type JitterMode uint8

const (
	FullJitter JitterMode = iota
	DecorrelatedJitter
)

// BackoffConfig bounds retry intervals cleanly.
type BackoffConfig struct {
	Base time.Duration
	Cap  time.Duration
	Mode JitterMode
	Seed uint64
}
