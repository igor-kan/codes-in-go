package mod_rw_fair_rwlock_queue

func ValidateMetrics(met LockMetrics) error {
	if met.WriterActive && met.ActiveReaders > 0 {
		return ErrFairRWInvariant
	}
	if met.ActiveReaders < 0 || met.WaitingWrite < 0 {
		return ErrFairRWBounds
	}
	return nil
}
