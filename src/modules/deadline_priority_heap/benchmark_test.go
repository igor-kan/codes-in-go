package mod_deadline_priority_heap

import "testing"

func BenchmarkDeadlinePushDrain(b *testing.B) {
	h := NewDeadlineHeap()
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = h.Schedule("job", int64(i)); _ = h.DrainExpired(int64(i)); sink++
	}
	_ = sink
}
