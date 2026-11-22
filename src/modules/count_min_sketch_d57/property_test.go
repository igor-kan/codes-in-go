package mod_count_min_sketch_d57

import "testing"

func TestProp_Cms(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	cm := NewCountMin(128)
	for k := uint64(1); k <= 20; k++ {
		cm.Add(k, k)
		if cm.Query(k) < k {
			t.Fatalf("underestimate invariant broken")
		}
	}
}
