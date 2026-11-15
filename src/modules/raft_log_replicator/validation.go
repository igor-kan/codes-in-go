package mod_raft_log_replicator

// ValidateRaftLog verifies sentinel, contiguous indices, and commit bounds.
func ValidateRaftLog(rl *RaftLog) error {
	if rl == nil || len(rl.Entries) == 0 || rl.Entries[0].Index != 0 {
		return ErrRaftInvalid
	}
	if rl.CommitIndex > rl.LastIndex() {
		return ErrRaftInvariant
	}
	for i, e := range rl.Entries {
		if e.Index != uint64(i) {
			return ErrRaftInvariant
		}
	}
	return nil
}
