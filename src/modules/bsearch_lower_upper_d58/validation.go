package mod_bsearch_lower_upper_d58

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrBSearchInvalid
	}
	return nil
}
