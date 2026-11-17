package mod_raft_election_timer_fsm

import "testing"

func BenchmarkRaftVoteCheck(b *testing.B) {
	n, _ := NewRaftNodeFSM(0, 5)
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		rep := n.HandleRequestVote(1, 1); if rep.VoteGranted { sink++ }
	}
	_ = sink
}
