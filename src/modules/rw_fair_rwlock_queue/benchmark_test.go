package mod_rw_fair_rwlock_queue

import "testing"

func BenchmarkFairRWRead(b *testing.B) {
	m := NewFairRWMutex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.RLock()
		m.RUnlock()
	}
}
