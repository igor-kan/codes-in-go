package mod_segment_tree_concurrent

type RangeStore interface {
	Update(idx int, val int64) error
	Query(l, r int) (SegStats, error)
}
