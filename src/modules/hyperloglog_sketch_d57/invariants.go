package mod_hyperloglog_sketch_d57

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrHllInvariant
	}
	return nil
}
