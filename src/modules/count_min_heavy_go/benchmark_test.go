package mod_count_min_heavy_go

import "testing"

func BenchmarkCmsAddEstimate(b *testing.B) {
	cms, _ := NewConservativeCMS(4, 256)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		cms.Add("event_a", 1); sink += int64(cms.Estimate("event_a"))
	}
	_ = sink
}
