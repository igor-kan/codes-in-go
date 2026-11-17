package mod_segment_tree_concurrent

func validateSegRange(n, l, r int) error {
	if l < 0 || r > n || l >= r {
		return ErrSegTreeBounds
	}
	return nil
}
