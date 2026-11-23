package mod_dag_longest_crit_d58

import "testing"

func Benchmark_Crit(b *testing.B) {
	esB := []TaskEdge{{0, 1, 2}, {1, 2, 5}, {2, 3, 3}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_, _ = DagMakespan(4, esB)
	}
	_ = sink
}
