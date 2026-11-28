package mod_boruvka_mst_go_d61

type MstSolver interface {
	MinimumSpanningTree(n int, edges []WeightedEdge) MstSummary
}

type TaggedComponent interface {
	ModuleSlug() string
}
