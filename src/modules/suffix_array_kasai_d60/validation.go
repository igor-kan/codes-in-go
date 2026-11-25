package mod_suffix_array_kasai_d60

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSuffixIndexInvalid
	}
	return nil
}
