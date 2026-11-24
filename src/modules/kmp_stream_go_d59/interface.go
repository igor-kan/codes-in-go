package mod_kmp_stream_go_d59

type PatternMatcher interface {
	FindAll(text, pat string) []int
}

type TaggedComponent interface {
	ModuleSlug() string
}
