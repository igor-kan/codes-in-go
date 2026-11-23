package mod_dag_longest_crit_d58

import "testing"

func TestProp_Crit(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	if _, ok := DagMakespan(2, []TaskEdge{{0, 1, 1}, {1, 0, 1}}); ok {
		t.Fatalf("expected cycle rejection")
	}
}
