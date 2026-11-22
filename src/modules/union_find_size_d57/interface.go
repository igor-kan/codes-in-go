package mod_union_find_size_d57

type Partition interface {
	Union(a, b int) bool
	Find(x int) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
