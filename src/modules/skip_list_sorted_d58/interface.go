package mod_skip_list_sorted_d58

type OrderedIndex interface {
	CountRange(lo, hi int64) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
