package mod_fenwick_range_point_go

type PrefixCounter interface {
	Add(idx int, delta int64)
	KthSmallest(k int64) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
