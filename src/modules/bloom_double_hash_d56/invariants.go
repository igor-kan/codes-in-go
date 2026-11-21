package mod_bloom_double_hash_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBloomInvariant
	}
	return nil
}
