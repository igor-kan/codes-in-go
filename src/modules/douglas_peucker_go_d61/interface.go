package mod_douglas_peucker_go_d61

type PolylineSimplifier interface {
	Simplify(points []Point, epsilon float64) []Point
}

type TaggedComponent interface {
	ModuleSlug() string
}
