package mod_booth_min_rotation_go_d61

import "testing"

func Benchmark_BoothRotation(b *testing.B) {
	s := "cbaabacbacccbbacababanana"
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(LeastRotation(s))
	}
	_ = sink
}
