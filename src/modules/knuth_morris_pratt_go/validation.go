package mod_knuth_morris_pratt_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrKmpGoInvalid
	}
	return nil
}
