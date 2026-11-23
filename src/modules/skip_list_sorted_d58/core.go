package mod_skip_list_sorted_d58

import "sort"

type SortedIndex struct {
	items []int64
}

func NewSortedIndex(raw []int64) *SortedIndex {
	cp := append([]int64(nil), raw...)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	return &SortedIndex{items: cp}
}

func (s *SortedIndex) CountRange(lo, hi int64) int {
	if lo > hi {
		return 0
	}
	l := sort.Search(len(s.items), func(i int) bool { return s.items[i] >= lo })
	r := sort.Search(len(s.items), func(i int) bool { return s.items[i] > hi })
	return r - l
}
