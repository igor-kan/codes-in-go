package mod_suffix_array_kasai_d60

type TextIndexer interface {
	DistinctCount(s string) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
