package mod_bplus_leaf_cursor_go

type RangeIndex interface {
	Put(k, v int64)
	Scan(lo, hi int64) []KVEntry
}

type TaggedComponent interface {
	ModuleSlug() string
}
