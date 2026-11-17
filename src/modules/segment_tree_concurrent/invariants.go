package mod_segment_tree_concurrent

func verifySegMatchesNaive(st *ConcurrentSegTree, raw []int64, l, r int) error {
	got, err := st.Query(l, r)
	if err != nil {
		return err
	}
	exp := fastSliceStats(raw, l, r)
	if got != exp {
		return ErrSegTreeInvariant
	}
	return nil
}
