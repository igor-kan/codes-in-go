package mod_bsearch_lower_upper_d58

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBSearchInvariant
	}
	return nil
}
