package mod_consistent_hash_vnodes_d56

type Router interface {
	Lookup(key string) string
}

type TaggedComponent interface {
	ModuleSlug() string
}
