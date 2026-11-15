package mod_topological_task_dag

import "testing"

func TestCycleDetection(t *testing.T) {
	plan, err := BuildSchedule([]int64{1, 1, 1}, []TaskEdge{{0, 1}, {1, 2}, {2, 0}})
	if err != nil || plan.IsDAG {
		t.Fatal("expected cycle detection")
	}
}
