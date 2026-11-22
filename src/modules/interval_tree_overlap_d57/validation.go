package mod_interval_tree_overlap_d57

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSweepInvalid
	}
	return nil
}
