package mod_consistent_ring_go_d60

type ShardRouter interface {
	RouteKey(key string) (string, bool)
}

type TaggedComponent interface {
	ModuleSlug() string
}
