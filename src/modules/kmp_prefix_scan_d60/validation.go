package mod_kmp_prefix_scan_d60

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrKmpMatcherInvalid
	}
	return nil
}
