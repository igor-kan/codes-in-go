package mod_hedged_request_quorum_go

import "testing"

func Benchmark_HedgedQuorum(b *testing.B) {
	reps := []ReplicaReply{{5, true}, {7, true}}
	b.ResetTimer()
	var sink int64
	for i := 0; i < b.N; i++ {
		v, _ := GatherQuorumSum(reps, 2); sink += v
	}
	_ = sink
}
