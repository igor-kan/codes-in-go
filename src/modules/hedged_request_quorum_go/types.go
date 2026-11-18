package mod_hedged_request_quorum_go

type ReplicaReply struct {
	Value int64
	Ok    bool
}

type ModuleMeta struct {
	Slug    string
	Version int
}
