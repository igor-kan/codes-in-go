package mod_union_find_rollback_d60

import "testing"

func Benchmark_RollbackDsu(b *testing.B) {
	d := NewDsuWithRollback(16)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		snap := d.Snapshot()
	d.Union(i&15, (i+1)&15)
	sink += int64(d.Components())
	d.Rollback(snap)
	}
	_ = sink
}
