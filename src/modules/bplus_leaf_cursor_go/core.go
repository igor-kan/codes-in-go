package mod_bplus_leaf_cursor_go

import "sort"

type BPlusLeafIndex struct {
	items []KVEntry
}

func NewBPlusLeafIndex() *BPlusLeafIndex {
	return &BPlusLeafIndex{}
}

func (idx *BPlusLeafIndex) Put(k, v int64) {
	i := sort.Search(len(idx.items), func(i int) bool { return idx.items[i].Key >= k })
	if i < len(idx.items) && idx.items[i].Key == k {
		idx.items[i].Val = v
		return
	}
	idx.items = append(idx.items, KVEntry{})
	copy(idx.items[i+1:], idx.items[i:])
	idx.items[i] = KVEntry{Key: k, Val: v}
}

func (idx *BPlusLeafIndex) Scan(lo, hi int64) []KVEntry {
	i := sort.Search(len(idx.items), func(i int) bool { return idx.items[i].Key >= lo })
	var out []KVEntry
	for ; i < len(idx.items) && idx.items[i].Key <= hi; i++ {
		out = append(out, idx.items[i])
	}
	return out
}
