package mod_brent_cycle_detect_go_d61

type CycleDetector interface {
	Detect(next func(int) int, start int) CycleInfo
}

type TaggedComponent interface {
	ModuleSlug() string
}
