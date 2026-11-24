package mod_min_heap_scheduler_d59

type JobQueue interface {
	Len() int
}

type TaggedComponent interface {
	ModuleSlug() string
}
