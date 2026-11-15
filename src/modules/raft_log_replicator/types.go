package mod_raft_log_replicator

// LogEntry stores a deterministic command stamped with leader Term.
type LogEntry struct {
	Index   uint64
	Term    uint64
	Command string
}

// AppendReq models a Raft AppendEntries RPC payload (1-based log indexing).
type AppendReq struct {
	Term         uint64
	PrevLogIndex uint64
	PrevLogTerm  uint64
	Entries      []LogEntry
	LeaderCommit uint64
}

// AppendResp returns follower acknowledgment and highest matched index.
type AppendResp struct {
	Term       uint64
	Success    bool
	MatchIndex uint64
}

// RaftLog tracks committed and replicated log entries on a node.
type RaftLog struct {
	CurrentTerm uint64
	CommitIndex uint64
	Entries     []LogEntry // sentinel dummy entry at index 0: {Index:0, Term:0}
}
