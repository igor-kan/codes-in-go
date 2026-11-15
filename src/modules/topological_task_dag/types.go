package mod_topological_task_dag

// TaskEdge directs dependency from U (predecessor) to V (successor).
type TaskEdge struct {
	U int
	V int
}

// SchedulePlan captures critical path timing metrics and topological waves.
type SchedulePlan struct {
	IsDAG      bool
	Makespan   int64
	EarlyStart []int64
	LateStart  []int64
	Waves      [][]int
}
