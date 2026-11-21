package mod_merkle_audit_tree_d56

import "testing"

func Benchmark_Merkle(b *testing.B) {
	mtB := NewMerkleTree([]uint64{11, 22, 33, 44, 55, 66, 77, 88})
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		_ = VerifyProof(33, mtB.Proof(2), mtB.Root())
	}
	_ = sink
}
