package mod_raft_log_replicator

// LastIndex returns the highest log index present.
func (rl *RaftLog) LastIndex() uint64 {
	return uint64(len(rl.Entries) - 1)
}

// TermAt returns the term at 1-based log idx, or (0, false) if missing.
func (rl *RaftLog) TermAt(idx uint64) (uint64, bool) {
	if idx >= uint64(len(rl.Entries)) {
		return 0, false
	}
	return rl.Entries[idx].Term, true
}

// HandleAppendEntries applies Raft Section 5.3 log-matching & conflict resolution.
func (rl *RaftLog) HandleAppendEntries(req AppendReq) AppendResp {
	if req.Term < rl.CurrentTerm {
		return AppendResp{Term: rl.CurrentTerm, Success: false}
	}
	if req.Term > rl.CurrentTerm {
		rl.CurrentTerm = req.Term
	}
	prevTerm, ok := rl.TermAt(req.PrevLogIndex)
	if !ok || prevTerm != req.PrevLogTerm {
		return AppendResp{Term: rl.CurrentTerm, Success: false}
	}
	for i, ent := range req.Entries {
		idx := req.PrevLogIndex + 1 + uint64(i)
		if idx <= rl.LastIndex() {
			if rl.Entries[idx].Term != ent.Term {
				rl.Entries = rl.Entries[:idx]
				rl.Entries = append(rl.Entries, req.Entries[i:]...)
				break
			}
		} else {
			rl.Entries = append(rl.Entries, req.Entries[i:]...)
			break
		}
	}
	if req.LeaderCommit > rl.CommitIndex {
		lastNew := rl.LastIndex()
		if req.LeaderCommit < lastNew {
			rl.CommitIndex = req.LeaderCommit
		} else {
			rl.CommitIndex = lastNew
		}
	}
	return AppendResp{Term: rl.CurrentTerm, Success: true, MatchIndex: rl.LastIndex()}
}
