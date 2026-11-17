package mod_deadline_priority_heap

func fastCountDueBefore(deadlines []int64, cutoff int64) int {
	cnt := 0
	for _, dl := range deadlines {
		if dl <= cutoff {
			cnt++
		}
	}
	return cnt
}
