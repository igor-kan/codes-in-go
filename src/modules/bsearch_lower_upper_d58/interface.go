package mod_bsearch_lower_upper_d58

type Finder interface {
	CountKey(xs []int64, target int64) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
