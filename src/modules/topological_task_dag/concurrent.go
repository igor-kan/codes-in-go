package mod_topological_task_dag

import "sync"

// ExecuteWaves executes each topological wave concurrently while respecting wave barriers.
func ExecuteWaves(plan *SchedulePlan, fn func(taskID int)) int {
	ran := 0
	for _, w := range plan.Waves {
		var wg sync.WaitGroup
		for _, tid := range w {
			idCopy := tid
			ran++
			wg.Add(1)
			go func() {
				defer wg.Done()
				fn(idCopy)
			}()
		}
		wg.Wait()
	}
	return ran
}
