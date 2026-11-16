package mod_b_tree_order_m

// OrderedDict exposes ordered insertion, point lookup, and sorted traversal.
type OrderedDict interface {
	Put(key, val int64)
	Get(key int64) (int64, bool)
	InOrder() []KVItem
}
