package mod_quadtree_spatial_index

func validateRect(r Rect) error {
	if r.MinX >= r.MaxX || r.MinY >= r.MaxY {
		return ErrQuadtreeInvalid
	}
	return nil
}
