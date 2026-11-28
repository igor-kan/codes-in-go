package mod_myers_bit_edit_go_d61

import "testing"

func Benchmark_MyersEdit(b *testing.B) {
	pat := "kitten"
	txt := "sitting on a keyboard"
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(MyersDistance(pat, txt))
	}
	_ = sink
}
