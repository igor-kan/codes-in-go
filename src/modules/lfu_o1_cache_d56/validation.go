package mod_lfu_o1_cache_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrLfuInvalid
	}
	return nil
}
