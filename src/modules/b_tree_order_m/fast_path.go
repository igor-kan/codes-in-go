package mod_b_tree_order_m

func (t *BTree) RangeScan(lo, hi int64) []KVItem {
	all := t.InOrder()
	out := make([]KVItem, 0)
	for _, it := range all {
		if it.Key >= lo && it.Key <= hi {
			out = append(out, it)
		}
	}
	return out
}
