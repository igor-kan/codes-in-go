package mod_bsearch_lower_upper_d58

func LowerBoundInt64(xs []int64, target int64) int {
	lo, hi := 0, len(xs)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if xs[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func UpperBoundInt64(xs []int64, target int64) int {
	lo, hi := 0, len(xs)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if xs[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func CountSortedOccurrences(xs []int64, target int64) int {
	return UpperBoundInt64(xs, target) - LowerBoundInt64(xs, target)
}
