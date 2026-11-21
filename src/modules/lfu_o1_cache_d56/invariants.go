package mod_lfu_o1_cache_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrLfuInvariant
	}
	return nil
}
