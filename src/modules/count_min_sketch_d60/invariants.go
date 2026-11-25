package mod_count_min_sketch_d60

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrCMSketchInvariant
	}
	return nil
}
