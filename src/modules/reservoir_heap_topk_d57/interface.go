package mod_reservoir_heap_topk_d57

type TopKCollector interface {
	Push(v int64)
	TopSorted() []int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
