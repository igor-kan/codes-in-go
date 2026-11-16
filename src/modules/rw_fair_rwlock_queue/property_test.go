package mod_rw_fair_rwlock_queue

import "testing"

func TestTicketOrderUncontended(t *testing.T) {
	m := NewFairRWMutex()
	for i := 0; i < 50; i++ {
		m.Lock()
		m.Unlock()
		m.RLock()
		m.RUnlock()
	}
	if err := ValidateMetrics(m.Snapshot()); err != nil {
		t.Fatal(err)
	}
}
