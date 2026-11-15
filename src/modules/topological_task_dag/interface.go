package mod_topological_task_dag

// BuildSchedule computes topological waves and Critical Path timings over tasks 0..n-1.
func BuildSchedule(durations []int64, edges []TaskEdge) (*SchedulePlan, error) {
	if len(durations) == 0 {
		return nil, ErrDAGEmpty
	}
	for _, e := range edges {
		if e.U < 0 || e.V < 0 || e.U >= len(durations) || e.V >= len(durations) {
			return nil, ErrDAGBounds
		}
	}
	return computeCPM(durations, edges), nil
}
