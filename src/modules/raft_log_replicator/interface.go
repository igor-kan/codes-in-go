package mod_raft_log_replicator

// NewRaftLog initializes a Raft log with sentinel entry at index 0.
func NewRaftLog() *RaftLog {
	return &RaftLog{
		Entries: []LogEntry{{Index: 0, Term: 0}},
	}
}
