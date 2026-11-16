package mod_b_tree_order_m

// KVItem stores an ordered integer key and payload value inside B-Tree nodes.
type KVItem struct {
	Key int64
	Val int64
}

type bnode struct {
	leaf     bool
	items    []KVItem
	children []*bnode
}
