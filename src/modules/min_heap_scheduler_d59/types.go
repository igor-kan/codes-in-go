package mod_min_heap_scheduler_d59

type TaskItem struct {
	Deadline int64
	Seq      int
	Payload  int
}

type ModuleMeta struct {
	Slug    string
	Version int
}
