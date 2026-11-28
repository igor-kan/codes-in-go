package mod_fenwick_2d_rect_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrFenwick2DInvalid
	}
	return nil
}
