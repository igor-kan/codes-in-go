package mod_count_min_sketch_d57

type FreqSketch interface {
	Add(item, count uint64)
	Query(item uint64) uint64
}

type TaggedComponent interface {
	ModuleSlug() string
}
