package mod_raft_election_timer_fsm

import "errors"

var (
	ErrRaftFsmInvalid   = errors.New("raftfsm: invalid argument")
	ErrRaftFsmBounds    = errors.New("raftfsm: index out of bounds")
	ErrRaftFsmEmpty     = errors.New("raftfsm: empty container")
	ErrRaftFsmInvariant = errors.New("raftfsm: invariant violated")
)
