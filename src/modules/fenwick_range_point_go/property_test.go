package mod_fenwick_range_point_go

import "testing"

func TestProp_BitGo(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	f := NewFenwickTree(8); f.Add(1, 5); f.Add(4, 2); if f.PrefixSum(4) != 7 { t.Fatal("bad prefix") }
}
