package mod_segment_tree_concurrent

type SegStats struct {
	Sum   int64
	Max   int64
	Count int
}

func (s SegStats) Average() float64 {
	if s.Count == 0 {
		return 0
	}
	return float64(s.Sum) / float64(s.Count)
}
