package mod_dinic_blocking_flow_go

type MaxFlowSolver interface {
	AddEdge(u, v int, cap int64)
	MaxFlow(s, t int) int64
}

type TaggedComponent interface {
	ModuleSlug() string
}
