package mod_raft_election_timer_fsm

type Role int

const (
	Follower Role = iota
	Candidate
	Leader
)

type VoteReply struct {
	Term        uint64
	VoteGranted bool
}
