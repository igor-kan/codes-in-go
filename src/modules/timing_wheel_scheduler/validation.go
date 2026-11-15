package mod_timing_wheel_scheduler

// ValidateTimingWheel checks bucket array integrity.
func ValidateTimingWheel(tw *TimingWheel) error {
	if tw == nil || tw.NumSlots <= 0 || len(tw.Buckets) != tw.NumSlots {
		return ErrTimerInvalid
	}
	return nil
}
