package mod_raft_election_timer_fsm

import "testing"

func TestRaftHigherTermPreemption(t *testing.T) {
	n, _ := NewRaftNodeFSM(0, 3)
	t1 := n.StartElection()
	if err := n.ReceiveHeartbeat(t1 + 2); err != nil {
		t.Fatal(err)
	}
	role, term, _ := n.Snapshot()
	if role != Follower || term != t1+2 {
		t.Fatalf("expected step down to Follower at higher term")
	}
}
