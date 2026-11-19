package mod_disjoint_set_union_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrDsuGoInvariant
	}
	return nil
}
