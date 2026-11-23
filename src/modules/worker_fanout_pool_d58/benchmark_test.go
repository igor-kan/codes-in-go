package mod_worker_fanout_pool_d58

import "testing"

func Benchmark_Pool(b *testing.B) {
	inB := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = ParallelSquareMap(inB, 2)
	}
	_ = sink
}
