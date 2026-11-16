package mod_rw_fair_rwlock_queue

func (m *FairRWMutex) Snapshot() LockMetrics {
	m.mu.Lock()
	out := LockMetrics{
		ActiveReaders: m.activeReaders,
		WriterActive:  m.writerActive,
		WaitingWrite:  m.waitingWrites,
	}
	m.mu.Unlock()
	return out
}
