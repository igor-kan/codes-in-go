package mod_cuckoo_filter_fp8_go

type DeletableFilter interface {
	Insert(key uint64) bool
	Lookup(key uint64) bool
	Delete(key uint64) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
