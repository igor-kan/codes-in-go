package mod_min_heap_scheduler_d59

import "testing"

func Benchmark_Sched(b *testing.B) {
	jobsB := []TaskItem{{50, 1, 1}, {10, 2, 2}, {30, 3, 3}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = RunByDeadline(jobsB)
	}
	_ = sink
}
