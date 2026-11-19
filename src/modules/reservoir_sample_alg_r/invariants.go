package mod_reservoir_sample_alg_r

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrReservoirGoInvariant
	}
	return nil
}
