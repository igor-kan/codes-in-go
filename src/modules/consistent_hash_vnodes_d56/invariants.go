package mod_consistent_hash_vnodes_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrRingInvariant
	}
	return nil
}
