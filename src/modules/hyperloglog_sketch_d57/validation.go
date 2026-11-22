package mod_hyperloglog_sketch_d57

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrHllInvalid
	}
	return nil
}
