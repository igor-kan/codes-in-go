package mod_knuth_morris_pratt_go

type PatternScanner interface {
	Search(text, pat string) []int
}

type TaggedComponent interface {
	ModuleSlug() string
}
