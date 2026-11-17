package mod_quadtree_spatial_index

type Point struct {
	X, Y float64
	ID   int
}

type Rect struct {
	MinX, MinY, MaxX, MaxY float64
}
