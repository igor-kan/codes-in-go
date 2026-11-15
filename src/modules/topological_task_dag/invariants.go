package mod_topological_task_dag

// VerifyCPMPrecendence asserts ES[v] >= ES[u] + duration[u] across all edges.
func VerifyCPMPrecendence(durations []int64, edges []TaskEdge, plan *SchedulePlan) error {
	if !plan.IsDAG {
		return ErrDAGInvariant
	}
	for _, e := range edges {
		if plan.EarlyStart[e.V] < plan.EarlyStart[e.U]+durations[e.U] {
			return ErrDAGInvariant
		}
	}
	return nil
}
