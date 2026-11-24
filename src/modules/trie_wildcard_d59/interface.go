package mod_trie_wildcard_d59

type WildMatcher interface {
	Search(pat string) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
