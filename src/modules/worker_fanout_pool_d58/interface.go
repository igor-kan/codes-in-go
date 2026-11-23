package mod_worker_fanout_pool_d58

type BatchMapper interface {
	MapInts(xs []int64) []int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
