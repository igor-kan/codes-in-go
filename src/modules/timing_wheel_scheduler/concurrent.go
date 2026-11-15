package mod_timing_wheel_scheduler

import "sync"

// LockedWheel guards TimingWheel mutations across concurrent worker ticks.
type LockedWheel struct {
	mu sync.Mutex
	w  *TimingWheel
}

// AdvanceLocked steps ticks under lock.
func (lw *LockedWheel) AdvanceLocked(ticks uint64) []*TimerEntry {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	return lw.w.Advance(ticks)
}
