package mod_timing_wheel_scheduler

// VerifyTimingWheelBuckets checks active map equals sum of slot list lengths.
func VerifyTimingWheelBuckets(tw *TimingWheel) error {
	if err := ValidateTimingWheel(tw); err != nil {
		return err
	}
	cnt := 0
	for _, b := range tw.Buckets {
		cnt += b.Len()
	}
	if cnt != len(tw.Active) {
		return ErrTimerInvariant
	}
	return nil
}
