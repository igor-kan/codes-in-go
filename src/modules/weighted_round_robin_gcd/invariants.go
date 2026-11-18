package mod_weighted_round_robin_gcd

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSmoothWrrInvariant
	}
	return nil
}
