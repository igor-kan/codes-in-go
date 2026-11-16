package mod_singleflight_dedup_group

import "testing"

func BenchmarkDedupSequential(b *testing.B) {
	g := NewDedupGroup()
	fn := func() (int64, error) { return 77, nil }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Do("k", fn)
	}
}
