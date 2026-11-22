package mod_interval_tree_overlap_d57

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSweepInvariant
	}
	return nil
}
