package mod_disjoint_set_union_go

type PartitionSet interface {
	Find(x int) int
	Union(a, b int) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
