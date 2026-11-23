package mod_worker_fanout_pool_d58

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrPoolInvariant
	}
	return nil
}
