package mod_jump_consistent_hash_go

import "testing"

func Benchmark_JumpHash(b *testing.B) {
	
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		sink += int64(JumpConsistentHash(uint64(i), 64))
	}
	_ = sink
}
