package mod_union_find_size_d57

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrDsuSizeInvalid
	}
	return nil
}
