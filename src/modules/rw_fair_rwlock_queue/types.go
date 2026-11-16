package mod_rw_fair_rwlock_queue

// LockMetrics snapshots contention state safely under monitor protection.
type LockMetrics struct {
	ActiveReaders int
	WriterActive  bool
	WaitingWrite  int
}
