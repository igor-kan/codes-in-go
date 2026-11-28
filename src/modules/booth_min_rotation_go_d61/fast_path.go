package mod_booth_min_rotation_go_d61

func fastFoldSlice(xs []int64) int64 {
	var h int64 = 1469598103934665603
	for _, v := range xs {
		h = (h ^ v) * 1099511628211
	}
	return h
}
