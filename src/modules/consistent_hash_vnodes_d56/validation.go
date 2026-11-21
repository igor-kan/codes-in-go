package mod_consistent_hash_vnodes_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrRingInvalid
	}
	return nil
}
