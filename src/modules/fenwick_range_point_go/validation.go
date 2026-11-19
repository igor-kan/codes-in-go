package mod_fenwick_range_point_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBitGoInvalid
	}
	return nil
}
