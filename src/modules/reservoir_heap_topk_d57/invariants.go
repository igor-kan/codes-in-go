package mod_reservoir_heap_topk_d57

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrTopKInvariant
	}
	return nil
}
