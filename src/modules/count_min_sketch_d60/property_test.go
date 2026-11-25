package mod_count_min_sketch_d60

import "testing"

func TestProp_CMSketch(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	cm := NewCountMinSketch(4, 64)
	exact := map[uint64]int64{}
	for k := uint64(1); k <= 40; k++ {
		cm.Add(k, int64(k%5+1))
		exact[k] += int64(k%5 + 1)
	}
	for k, ex := range exact {
		if cm.Estimate(k) < ex {
			t.Fatal("underestimation violates Count-Min guarantee")
		}
	}
}
