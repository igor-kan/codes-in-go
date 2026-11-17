package mod_raft_election_timer_fsm

func validateClusterSize(id, peers int) error {
	if peers < 1 || id < 0 || id >= peers {
		return ErrRaftFsmBounds
	}
	return nil
}
