package mod_sliding_percentile_p2_go

import "testing"

func TestProp_WindowQuant(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	w := NewSlidingQuantileWindow(4)
	w.Push(2); w.Push(4); w.Push(6); w.Push(8)
	if w.Summary().Median != 5 {
		t.Fatal("expected median 5")
	}
}
