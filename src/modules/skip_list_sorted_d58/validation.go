package mod_skip_list_sorted_d58

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSkipChunkInvalid
	}
	return nil
}
