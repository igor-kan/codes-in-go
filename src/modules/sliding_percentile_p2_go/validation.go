package mod_sliding_percentile_p2_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrWindowQuantInvalid
	}
	return nil
}
