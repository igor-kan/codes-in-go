package mod_hedged_request_quorum_go

func GatherQuorumSum(replicas []ReplicaReply, k int) (int64, error) {
	if k <= 0 || k > len(replicas) {
		return 0, ErrHedgedQuorumInvalid
	}
	ch := make(chan ReplicaReply, len(replicas))
	for _, rep := range replicas {
		r := rep
		go func() { ch <- r }()
	}
	var sum int64
	okCnt := 0
	for i := 0; i < len(replicas); i++ {
		msg := <-ch
		if msg.Ok {
			sum += msg.Value
			okCnt++
			if okCnt == k {
				return sum, nil
			}
		}
	}
	return 0, ErrHedgedQuorumInvariant
}
