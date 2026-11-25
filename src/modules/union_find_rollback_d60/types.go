package mod_union_find_rollback_d60

type UndoEntry struct {
	Child int
	Root  int
	Rank  int
}

type ModuleMeta struct {
	Slug    string
	Version int
}
