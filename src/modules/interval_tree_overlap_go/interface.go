package mod_interval_tree_overlap_go

type IntervalQueryEngine interface {
	QueryOverlap(lo, hi int64) []ClosedIv
}

type TaggedComponent interface {
	ModuleSlug() string
}
