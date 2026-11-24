package mod_dijkstra_sp_go_d59

type ShortestSolver interface {
	Distances(n, src int, adj [][]WeightedArc) []int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
