package mod_sliding_log_limiter_d57

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSLogInvalid
	}
	return nil
}
