package mod_sliding_log_limiter_d57

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSLogInvariant
	}
	return nil
}
