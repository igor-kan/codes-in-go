package mod_raft_election_timer_fsm

func fastQuorum(peers int) int {
	return peers/2 + 1
}

func hasMajorityVotes(votes, peers int) bool {
	return votes >= fastQuorum(peers)
}
