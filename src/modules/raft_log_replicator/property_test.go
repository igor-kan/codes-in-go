package mod_raft_log_replicator

import "testing"

func TestRaftIdempotentReDelivery(t *testing.T) {
	node := NewRaftLog()
	req := AppendReq{
		Term:         3,
		PrevLogIndex: 0,
		PrevLogTerm:  0,
		Entries:      []LogEntry{{Index: 1, Term: 3, Command: "a"}, {Index: 2, Term: 3, Command: "b"}},
		LeaderCommit: 1,
	}
	_ = node.HandleAppendEntries(req)
	snapLen := node.LastIndex()
	_ = node.HandleAppendEntries(req)
	if node.LastIndex() != snapLen || ValidateRaftLog(node) != nil {
		t.Fatal("idempotent replay mutated log length")
	}
}
