package mod_quadtree_spatial_index

type SpatialIndex interface {
	Insert(p Point) error
	QueryRange(box Rect) []Point
	Count() int
}
