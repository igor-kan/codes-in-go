package mod_reservoir_heap_topk_d57

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrTopKInvalid
	}
	return nil
}
