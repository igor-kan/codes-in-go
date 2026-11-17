package mod_count_min_heavy_go

func fastTotalMass(row []uint64) uint64 {
	var s uint64
	for _, v := range row {
		s += v
	}
	return s
}
