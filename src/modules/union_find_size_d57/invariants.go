package mod_union_find_size_d57

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrDsuSizeInvariant
	}
	return nil
}
