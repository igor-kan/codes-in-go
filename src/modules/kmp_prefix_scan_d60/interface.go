package mod_kmp_prefix_scan_d60

type PatternFinder interface {
	FindAll(text, pat string) []int
}

type TaggedComponent interface {
	ModuleSlug() string
}
