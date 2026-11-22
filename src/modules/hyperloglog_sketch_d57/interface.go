package mod_hyperloglog_sketch_d57

type DistinctSketch interface {
	AddHash(v uint64)
	Estimate() float64
}

type TaggedComponent interface {
	ModuleSlug() string
}
