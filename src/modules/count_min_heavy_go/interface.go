package mod_count_min_heavy_go

type FrequencySketch interface {
	Add(key string, delta uint64)
	Estimate(key string) uint64
}
