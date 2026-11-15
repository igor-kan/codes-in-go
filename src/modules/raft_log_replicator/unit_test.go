package mod_raft_log_replicator

import "testing"

func TestRaftAppendAndConflictTruncate(t *testing.T) {
	follower := NewRaftLog()
	follower.CurrentTerm = 1
	follower.AppendLocalBatch([]string{"set x=1", "stale"})
	// Leader in term 2 overwrites index 2 cleanly
	resp := follower.HandleAppendEntries(AppendReq{
		Term:         2,
		PrevLogIndex: 1,
		PrevLogTerm:  1,
		Entries:      []LogEntry{{Index: 2, Term: 2, Command: "set x=2"}, {Index: 3, Term: 2, Command: "set y=9"}},
		LeaderCommit: 2,
	})
	if !resp.Success || resp.MatchIndex != 3 || follower.CommitIndex != 2 {
		t.Fatalf("unexpected response: %+v commit=%d", resp, follower.CommitIndex)
	}
	if follower.Entries[2].Command != "set x=2" {
		t.Fatalf("conflict entry not overwritten: %+v", follower.Entries[2])
	}
	if err := ValidateRaftLog(follower); err != nil {
		t.Fatal(err)
	}
}
