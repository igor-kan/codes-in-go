package mod_count_min_sketch_d60

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrCMSketchInvalid
	}
	return nil
}
