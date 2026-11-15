package mod_topological_task_dag

// ValidateDurations ensures non-negative task execution times.
func ValidateDurations(durs []int64) error {
	if len(durs) == 0 {
		return ErrDAGEmpty
	}
	for _, d := range durs {
		if d < 0 {
			return ErrDAGInvalid
		}
	}
	return nil
}
