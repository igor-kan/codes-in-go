package mod_timing_wheel_scheduler

import "errors"

var (
	ErrTimerEmpty     = errors.New("timer: empty or nil input")
	ErrTimerBounds    = errors.New("timer: index or dimension out of bounds")
	ErrTimerInvalid   = errors.New("timer: invalid configuration or state")
	ErrTimerInvariant = errors.New("timer: structural invariant violated")
)
