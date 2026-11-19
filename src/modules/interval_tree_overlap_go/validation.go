package mod_interval_tree_overlap_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrIvTreeGoInvalid
	}
	return nil
}
