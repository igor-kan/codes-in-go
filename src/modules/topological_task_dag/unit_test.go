package mod_topological_task_dag

import "testing"

func TestCPMCriticalPath(t *testing.T) {
	// Task 0 (dur 3) -> Task 1 (dur 4), Task 0 -> Task 2 (dur 1), Task 1 & 2 -> Task 3 (dur 2)
	durs := []int64{3, 4, 1, 2}
	edges := []TaskEdge{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	plan, err := BuildSchedule(durs, edges)
	if err != nil || !plan.IsDAG || plan.Makespan != 9 {
		t.Fatalf("bad makespan: %+v err=%v", plan, err)
	}
	if err := VerifyCPMPrecendence(durs, edges, plan); err != nil {
		t.Fatal(err)
	}
	crit := CriticalTasks(plan)
	if len(crit) != 3 || crit[0] != 0 || crit[1] != 1 || crit[2] != 3 {
		t.Fatalf("unexpected critical path: %v", crit)
	}
}
