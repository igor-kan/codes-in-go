package mod_token_bucket_limiter_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrLimiterInvalid
	}
	return nil
}
