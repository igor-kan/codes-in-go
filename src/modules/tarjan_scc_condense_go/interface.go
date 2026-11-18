package mod_tarjan_scc_condense_go

type SccFinder interface {
	Compute(n int, edges [][2]int) SccPartition
}

type TaggedComponent interface {
	ModuleSlug() string
}
