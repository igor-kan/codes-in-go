package mod_rw_fair_rwlock_queue

func VerifyMutualExclusion(writers, ops int, observed int64) error {
	if observed != int64(writers*ops) {
		return ErrFairRWInvariant
	}
	return nil
}
