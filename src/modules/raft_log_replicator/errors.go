package mod_raft_log_replicator

import "errors"

var (
	ErrRaftEmpty     = errors.New("raft: empty or nil input")
	ErrRaftBounds    = errors.New("raft: index or dimension out of bounds")
	ErrRaftInvalid   = errors.New("raft: invalid configuration or state")
	ErrRaftInvariant = errors.New("raft: structural invariant violated")
)
