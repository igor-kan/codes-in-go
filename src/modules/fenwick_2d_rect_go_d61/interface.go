package mod_fenwick_2d_rect_go_d61

type GridRangeSummer interface {
	Add(x, y int, delta int64)
	RectSum(q RectQuery) int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
