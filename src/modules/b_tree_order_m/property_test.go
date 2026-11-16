package mod_b_tree_order_m

import "testing"

func TestBTreeOracleParity(t *testing.T) {
	bt, _ := NewBTree(2)
	ref := make(map[int64]int64)
	var keys []int64
	for i := int64(0); i < 200; i++ {
		k := (i*47 + 13) % 151
		bt.Put(k, i*3)
		ref[k] = i * 3
		keys = append(keys, k)
	}
	if err := VerifyBTreeInvariants(bt); err != nil {
		t.Fatal(err)
	}
	found := ParallelBTreeLookup(bt, keys)
	for i, k := range keys {
		v, ok := bt.Get(k)
		if !ok || !found[i] || v != ref[k] {
			t.Fatalf("mismatch at k=%d: got=%d want=%d", k, v, ref[k])
		}
	}
}
