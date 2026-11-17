package mod_raft_election_timer_fsm

func verifySingleLeaderInTerm(nodes []*RaftNodeFSM, term uint64) error {
	leaders := 0
	for _, n := range nodes {
		role, t, _ := n.Snapshot()
		if role == Leader && t == term {
			leaders++
		}
	}
	if leaders > 1 {
		return ErrRaftFsmInvariant
	}
	return nil
}
