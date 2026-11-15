package mod_raft_log_replicator

import "testing"

func BenchmarkRaftAppendBatch(b *testing.B) {
	n := NewRaftLog()
	n.CurrentTerm = 1
	for i := 0; i < b.N; i++ {
		idx := n.LastIndex() + 1
		n.HandleAppendEntries(AppendReq{
			Term:         1,
			PrevLogIndex: idx - 1,
			PrevLogTerm:  1,
			Entries:      []LogEntry{{Index: idx, Term: 1, Command: "noop"}},
			LeaderCommit: idx,
		})
	}
}
