package mod_bloom_double_hash_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBloomInvalid
	}
	return nil
}
