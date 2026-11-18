package mod_weighted_round_robin_gcd

import "testing"

func TestProp_SmoothWrr(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	b := NewSmoothWRR([]int{2, 2})
	if b.NextPeer() == b.NextPeer() {
		t.Fatal("equal weights must alternate")
	}
}
