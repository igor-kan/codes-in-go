package mod_eertree_palindrome_go_d61

import "testing"

func Benchmark_Eertree(b *testing.B) {
	sample := "abacabadabacabaeabacabadabacaba"
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(DistinctPalindromes(sample))
	}
	_ = sink
}
