package mod_circuit_halfopen_d58

type BreakerPolicy struct {
	FailureLimit int
	CooldownMs   int64
}

type ModuleMeta struct {
	Slug    string
	Version int
}
