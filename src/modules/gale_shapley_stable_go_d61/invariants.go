package mod_gale_shapley_stable_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrStableMarriageInvariant
	}
	return nil
}
