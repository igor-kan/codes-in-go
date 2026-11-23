package mod_circuit_halfopen_d58

import "testing"

func TestProp_Breaker(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	cb := NewBreaker(3, 100)
	cb.OnFail(0)
	if !cb.Allow(5) {
		t.Fatalf("opened prematurely")
	}
}
