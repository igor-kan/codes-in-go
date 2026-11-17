package mod_quadtree_spatial_index

func fastBoxContains(box Rect, pts []Point) []Point {
	out := make([]Point, 0, 8)
	for _, p := range pts {
		if p.X >= box.MinX && p.X <= box.MaxX && p.Y >= box.MinY && p.Y <= box.MaxY {
			out = append(out, p)
		}
	}
	return out
}
