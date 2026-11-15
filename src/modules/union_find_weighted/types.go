package mod_union_find_weighted

// WeightedDSU tracks disjoint equivalence classes with additive potential differences.
type WeightedDSU struct {
	Parent []int
	Rank   []int
	Weight []float64 // weight[x] = val[x] - val[parent[x]]
}
