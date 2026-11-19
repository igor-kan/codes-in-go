package mod_interval_tree_overlap_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrIvTreeGoInvariant
	}
	return nil
}
