package mod_aho_corasick_go

type MultiMatcher interface {
	CountMatches(text []byte) int
	StateCount() int
}
