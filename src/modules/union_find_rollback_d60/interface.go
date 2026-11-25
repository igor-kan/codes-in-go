package mod_union_find_rollback_d60

type RewindableSet interface {
	Union(a, b int) bool
	Rollback(targetLen int)
}

type TaggedComponent interface {
	ModuleSlug() string
}
