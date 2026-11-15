package mod_topological_task_dag

import "testing"

func BenchmarkCPMDAG(b *testing.B) {
	durs := make([]int64, 200)
	edges := make([]TaskEdge, 0, 199)
	for i := 0; i < 199; i++ {
		durs[i] = 2
		edges = append(edges, TaskEdge{i, i + 1})
	}
	durs[199] = 2
	for i := 0; i < b.N; i++ {
		_, _ = BuildSchedule(durs, edges)
	}
}
