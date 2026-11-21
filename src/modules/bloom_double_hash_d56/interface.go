package mod_bloom_double_hash_d56

type ProbSet interface {
	Add(v uint64)
	MightContain(v uint64) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
