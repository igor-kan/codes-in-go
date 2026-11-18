package mod_sliding_percentile_p2_go

type PercentileWindow interface {
	Push(x float64)
	Summary() WindowSummary
}

type TaggedComponent interface {
	ModuleSlug() string
}
