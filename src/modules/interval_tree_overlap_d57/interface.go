package mod_interval_tree_overlap_d57

type OverlapMeasurer interface {
	MaxOverlap(spans []TimeSpan) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
