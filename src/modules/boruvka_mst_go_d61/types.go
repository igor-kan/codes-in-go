package mod_boruvka_mst_go_d61

type WeightedEdge struct {
	U int
	V int
	W int64
}

type MstSummary struct {
	Weight int64
	Edges  int
}

type ModuleMeta struct {
	Slug    string
	Version int
}
