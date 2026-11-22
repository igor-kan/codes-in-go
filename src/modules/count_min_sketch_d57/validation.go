package mod_count_min_sketch_d57

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrCmsInvalid
	}
	return nil
}
