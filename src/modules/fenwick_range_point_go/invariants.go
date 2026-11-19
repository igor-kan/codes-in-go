package mod_fenwick_range_point_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrBitGoInvariant
	}
	return nil
}
