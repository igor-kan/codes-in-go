package mod_dag_longest_crit_d58

type Scheduler interface {
	Makespan(n int, edges []TaskEdge) (int64, bool)
}

type TaggedComponent interface {
	ModuleSlug() string
}
