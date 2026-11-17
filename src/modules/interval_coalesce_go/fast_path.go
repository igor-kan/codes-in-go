package mod_interval_coalesce_go

func fastTotalCovered(spans []Span) int64 {
	var tot int64
	for _, s := range spans {
		tot += s.End - s.Start + 1
	}
	return tot
}
