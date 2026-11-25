package mod_suffix_array_kasai_d60

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSuffixIndexInvariant
	}
	return nil
}
