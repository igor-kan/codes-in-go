package mod_consistent_hash_vnodes_d56

import "testing"

func TestProp_Ring(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	hr := NewHashRing([]string{"n1", "n2"}, 8)
	h := hr.Lookup("probe-key")
	if h != "n1" && h != "n2" {
		t.Fatalf("unexpected host %s", h)
	}
}
