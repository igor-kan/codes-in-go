package mod_reservoir_sample_alg_r

type StreamSampler interface {
	Offer(val int64)
	Items() []int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
