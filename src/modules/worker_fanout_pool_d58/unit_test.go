package mod_worker_fanout_pool_d58

import "testing"

func TestUnit_Pool(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	got := ParallelSquareMap([]int64{1, -2, 3, 4, 5}, 3)
	if got[1] != 4 || got[4] != 25 {
		t.Fatalf("bad square map: %v", got)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
