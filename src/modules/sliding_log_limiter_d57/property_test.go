package mod_sliding_log_limiter_d57

import "testing"

func TestProp_SLog(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	sl := NewSlidingLog(3, 5)
	allowed := 0
	for ts := int64(0); ts < 5; ts++ {
		if sl.Allow(ts) {
			allowed++
		}
	}
	if allowed != 3 {
		t.Fatalf("expected 3 admitted in [0,4], got %d", allowed)
	}
}
