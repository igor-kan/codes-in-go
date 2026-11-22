package mod_count_min_sketch_d57

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrCmsInvariant
	}
	return nil
}
