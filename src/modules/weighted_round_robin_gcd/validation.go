package mod_weighted_round_robin_gcd

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSmoothWrrInvalid
	}
	return nil
}
