package mod_interval_coalesce_go

import "testing"

func BenchmarkIntervalContains(b *testing.B) {
	l := NewDisjointIntervals(); for k:=int64(0);k<64;k++{ _=l.Add(k*10, k*10+4) }
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		if l.Contains(int64(i % 640)) { sink++ }
	}
	_ = sink
}
