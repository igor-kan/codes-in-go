package mod_aho_multi_scan_d59

type KeywordScanner interface {
	CountHits(text string) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
