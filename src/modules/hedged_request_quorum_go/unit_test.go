package mod_hedged_request_quorum_go

import "testing"

func TestUnit_HedgedQuorum(t *testing.T) {
	if err := validatePositiveCap(1); err != nil {
		t.Fatal(err)
	}
	s, err := GatherQuorumSum([]ReplicaReply{{10, true}, {20, false}, {30, true}}, 2)
	if err != nil || s != 40 {
		t.Fatalf("expected 40, got %d err=%v", s, err)
	}
	if err := verifyPostcondition(true); err != nil {
		t.Fatal(err)
	}
}
