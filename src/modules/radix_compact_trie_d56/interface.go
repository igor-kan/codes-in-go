package mod_radix_compact_trie_d56

type PrefixIndex interface {
	Insert(word string)
	CountPrefix(prefix string) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
