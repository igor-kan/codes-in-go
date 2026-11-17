package mod_quadtree_spatial_index

func verifyQuadtreeParity(qt *Quadtree, all []Point, box Rect) error {
	got := qt.QueryRange(box)
	expected := fastBoxContains(box, all)
	if len(got) != len(expected) {
		return ErrQuadtreeInvariant
	}
	return nil
}
