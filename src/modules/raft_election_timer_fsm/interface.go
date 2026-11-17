package mod_raft_election_timer_fsm

type ElectionFSM interface {
	StartElection() uint64
	HandleVoteReply(from int, reply VoteReply) Role
	ReceiveHeartbeat(leaderTerm uint64) error
}
