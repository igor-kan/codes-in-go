package mod_cuckoo_filter_fp8_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrCuckooGoInvalid
	}
	return nil
}
