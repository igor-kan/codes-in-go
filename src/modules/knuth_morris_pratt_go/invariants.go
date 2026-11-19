package mod_knuth_morris_pratt_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrKmpGoInvariant
	}
	return nil
}
