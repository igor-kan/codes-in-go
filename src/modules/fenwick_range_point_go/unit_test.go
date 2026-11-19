package mod_fenwick_range_point_go

import "testing"

func TestUnit_BitGo(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	f := NewFenwickTree(16)
	f.Add(2, 3); f.Add(5, 2); f.Add(10, 4)
	if f.KthSmallest(1) != 2 || f.KthSmallest(4) != 5 || f.KthSmallest(7) != 10 {
		t.Fatal("unexpected kth smallest")
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
