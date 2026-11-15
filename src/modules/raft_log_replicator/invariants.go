package mod_raft_log_replicator

// VerifyLogMatching asserts two logs sharing (idx, term) match at every earlier index.
func VerifyLogMatching(a, b *RaftLog, idx uint64) error {
	ta, oka := a.TermAt(idx)
	tb, okb := b.TermAt(idx)
	if !oka || !okb || ta != tb {
		return nil
	}
	for k := uint64(1); k <= idx; k++ {
		if a.Entries[k] != b.Entries[k] {
			return ErrRaftInvariant
		}
	}
	return nil
}
