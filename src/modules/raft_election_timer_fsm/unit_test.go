package mod_raft_election_timer_fsm

import "testing"

func TestRaftElectionQuorumAndStepDown(t *testing.T) {
	nodes := make([]*RaftNodeFSM, 5)
	for i := range nodes {
		nodes[i], _ = NewRaftNodeFSM(i, 5)
	}
	term := nodes[0].StartElection()
	reps := broadcastVoteRequests(nodes, 0, term)
	var finalRole Role
	for i, rep := range reps {
		finalRole = nodes[0].HandleVoteReply(i, rep)
	}
	if finalRole != Leader {
		t.Fatalf("expected Leader after majority votes")
	}
	if err := verifySingleLeaderInTerm(nodes, term); err != nil {
		t.Fatal(err)
	}
	// Reject second candidate in same term since nodes voted for 0
	reps2 := broadcastVoteRequests(nodes[1:], 1, term)
	granted := 0
	for _, rep := range reps2 {
		if rep.VoteGranted {
			granted++
		}
	}
	if granted != 0 {
		t.Fatalf("expected 0 split votes for candidate 1 in term %d, got %d", term, granted)
	}
}
