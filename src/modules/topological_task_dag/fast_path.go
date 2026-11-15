package mod_topological_task_dag

// CriticalTasks returns indices with zero scheduling slack.
func CriticalTasks(plan *SchedulePlan) []int {
	var crit []int
	for i := range plan.EarlyStart {
		if plan.EarlyStart[i] == plan.LateStart[i] {
			crit = append(crit, i)
		}
	}
	return crit
}
