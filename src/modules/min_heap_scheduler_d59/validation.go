package mod_min_heap_scheduler_d59

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSchedInvalid
	}
	return nil
}
