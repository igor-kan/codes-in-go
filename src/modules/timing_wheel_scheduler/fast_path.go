package mod_timing_wheel_scheduler

// PendingCount returns active un-fired timer count.
func (tw *TimingWheel) PendingCount() int {
	return len(tw.Active)
}
