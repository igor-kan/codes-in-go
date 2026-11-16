package mod_rw_fair_rwlock_queue

// FairLocker enforces mutually exclusive writes and concurrent starvation-free reads.
type FairLocker interface {
	RLock()
	RUnlock()
	Lock()
	Unlock()
}
