package mod_count_min_heavy_go

type HeavyItem struct {
	Key   string
	Score uint64
	Rank  int
}

func (h HeavyItem) MeetsThreshold(minScore uint64) bool {
	return h.Score >= minScore
}
