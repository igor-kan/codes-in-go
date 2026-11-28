package mod_brent_cycle_detect_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBrentCycleInvalid
	}
	return nil
}
