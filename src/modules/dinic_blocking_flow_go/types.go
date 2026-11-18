package mod_dinic_blocking_flow_go

type DinicEdge struct {
	To  int
	Rev int
	Cap int64
}

type ModuleMeta struct {
	Slug    string
	Version int
}
