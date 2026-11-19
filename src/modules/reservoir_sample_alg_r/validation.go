package mod_reservoir_sample_alg_r

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrReservoirGoInvalid
	}
	return nil
}
