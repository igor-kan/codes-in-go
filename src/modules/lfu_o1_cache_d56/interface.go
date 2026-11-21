package mod_lfu_o1_cache_d56

type CacheBackend interface {
	Get(key int) (int, bool)
	Put(key, value int)
}

type TaggedComponent interface {
	ModuleSlug() string
}
