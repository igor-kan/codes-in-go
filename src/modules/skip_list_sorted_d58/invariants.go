package mod_skip_list_sorted_d58

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSkipChunkInvariant
	}
	return nil
}
