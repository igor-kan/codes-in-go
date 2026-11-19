package mod_disjoint_set_union_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrDsuGoInvalid
	}
	return nil
}
