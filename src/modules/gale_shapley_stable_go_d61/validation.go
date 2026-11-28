package mod_gale_shapley_stable_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrStableMarriageInvalid
	}
	return nil
}
