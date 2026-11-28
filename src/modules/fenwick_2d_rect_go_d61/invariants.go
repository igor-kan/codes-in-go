package mod_fenwick_2d_rect_go_d61

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrFenwick2DInvariant
	}
	return nil
}
