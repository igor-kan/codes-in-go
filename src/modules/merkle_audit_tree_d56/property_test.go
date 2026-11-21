package mod_merkle_audit_tree_d56

import "testing"

func TestProp_Merkle(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	mt := NewMerkleTree([]uint64{1, 2, 3, 4})
	for idx, v := range []uint64{1, 2, 3, 4} {
		if !VerifyProof(v, mt.Proof(idx), mt.Root()) {
			t.Fatalf("leaf %d proof failed", idx)
		}
	}
}
