package mod_hedged_request_quorum_go

import "testing"

func TestProp_HedgedQuorum(t *testing.T) {
	if parallelFoldSlices([]int64{1, 2}, []int64{3, 4}) == 0 {
		t.Fatal("unexpected zero fold")
	}
	_, err := GatherQuorumSum([]ReplicaReply{{10, false}, {20, false}}, 1)
	if err == nil {
		t.Fatal("expected quorum error")
	}
}
