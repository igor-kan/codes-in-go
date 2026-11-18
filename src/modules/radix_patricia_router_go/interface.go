package mod_radix_patricia_router_go

type PrefixRouter interface {
	Insert(prefix string, handler int)
	MatchLongest(path string) RouteMatch
}

type TaggedComponent interface {
	ModuleSlug() string
}
