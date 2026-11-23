package mod_worker_fanout_pool_d58

import "testing"

func TestProp_Pool(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	in := []int64{7, 8, 9, 10}
	a := ParallelSquareMap(in, 1)
	b := ParallelSquareMap(in, 4)
	for i := range in {
		if a[i] != b[i] {
			t.Fatalf("worker determinism failed")
		}
	}
}
