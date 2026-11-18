package mod_cuckoo_filter_fp8_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrCuckooGoInvariant
	}
	return nil
}
