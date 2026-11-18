package mod_jump_consistent_hash_go

import "testing"

func TestProp_JumpHash(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	if JumpConsistentHash(42, 1) != 0 {
		t.Fatal("single bucket must return 0")
	}
}
