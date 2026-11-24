package mod_min_heap_scheduler_d59

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSchedInvariant
	}
	return nil
}
