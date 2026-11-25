package mod_count_min_sketch_d60

type FrequencyEstimator interface {
	Add(key uint64, count int64)
	Estimate(key uint64) int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
