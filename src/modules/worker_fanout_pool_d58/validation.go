package mod_worker_fanout_pool_d58

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrPoolInvalid
	}
	return nil
}
