package mod_bplus_leaf_cursor_go

import "testing"

func Benchmark_BPlusCursor(b *testing.B) {
	idx := NewBPlusLeafIndex(); for k := int64(0); k < 32; k++ { idx.Put(k, k) }
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(len(idx.Scan(4, 20)))
	}
	_ = sink
}
