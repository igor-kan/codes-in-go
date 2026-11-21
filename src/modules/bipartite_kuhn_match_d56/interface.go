package mod_bipartite_kuhn_match_d56

type BipartiteSolver interface {
	MaxMatch() int
}

type TaggedComponent interface {
	ModuleSlug() string
}
