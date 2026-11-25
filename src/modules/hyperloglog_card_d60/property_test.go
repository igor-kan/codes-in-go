package mod_hyperloglog_card_d60

import "testing"

func TestProp_HllEstimator(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	h1 := NewHyperLogLog(6)
	h2 := NewHyperLogLog(6)
	for i := uint64(0); i < 50; i++ {
		h1.Observe(i * 13)
	}
	for i := uint64(0); i < 50; i++ {
		h2.Observe(i * 13)
		h2.Observe(i * 13)
	}
	if h1.Count() != h2.Count() {
		t.Fatal("duplicates changed HLL registers")
	}
}
