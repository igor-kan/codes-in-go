package mod_worker_fanout_pool_d58

type PoolSpec struct {
	Workers int
	Jobs    int
}

type ModuleMeta struct {
	Slug    string
	Version int
}
