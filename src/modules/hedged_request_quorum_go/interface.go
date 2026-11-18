package mod_hedged_request_quorum_go

type QuorumRunner interface {
	Gather(vals []int64, k int) (int64, error)
}

type TaggedComponent interface {
	ModuleSlug() string
}
