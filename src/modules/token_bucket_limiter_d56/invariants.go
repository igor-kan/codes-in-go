package mod_token_bucket_limiter_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrLimiterInvariant
	}
	return nil
}
