package mod_circuit_halfopen_d58

type CircuitGuard interface {
	Allow(nowMs int64) bool
}

type TaggedComponent interface {
	ModuleSlug() string
}
