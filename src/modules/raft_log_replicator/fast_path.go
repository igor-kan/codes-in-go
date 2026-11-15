package mod_raft_log_replicator

// AppendLocalBatch appends pre-validated sequential entries from an active leader.
func (rl *RaftLog) AppendLocalBatch(cmds []string) uint64 {
	for _, c := range cmds {
		idx := rl.LastIndex() + 1
		rl.Entries = append(rl.Entries, LogEntry{Index: idx, Term: rl.CurrentTerm, Command: c})
	}
	return rl.LastIndex()
}
