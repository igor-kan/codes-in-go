package mod_segment_tree_concurrent

import "math"

func fastSliceStats(arr []int64, l, r int) SegStats {
	var sum int64
	mx := int64(math.MinInt64)
	for i := l; i < r; i++ {
		sum += arr[i]
		if arr[i] > mx {
			mx = arr[i]
		}
	}
	return SegStats{Sum: sum, Max: mx}
}
