package mod_radix_patricia_router_go

type RouteMatch struct {
	Prefix  string
	Handler int
	Found   bool
}

type ModuleMeta struct {
	Slug    string
	Version int
}
