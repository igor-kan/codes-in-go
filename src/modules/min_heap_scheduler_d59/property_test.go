package mod_min_heap_scheduler_d59

import "testing"

func TestProp_Sched(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	items := []TaskItem{{5, 1, 9}, {1, 2, 4}, {3, 3, 7}}
	if got := RunByDeadline(items); got[0] != 4 || got[2] != 9 {
		t.Fatalf("bad sort: %v", got)
	}
}
