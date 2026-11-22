package mod_hyperloglog_sketch_d57

import "testing"

func TestProp_Hll(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	sk := NewHll64()
	for x := 0; x < 50; x++ {
		sk.Add(999)
	}
	if est := sk.Estimate(); est > 4.0 {
		t.Fatalf("duplicate flood inflated cardinality: %v", est)
	}
}
