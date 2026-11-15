package mod_raft_log_replicator

import "sync"

// ReplicatedQuorum checks if at least N/2+1 logs hold a matching entry at targetIdx.
func ReplicatedQuorum(logs []*RaftLog, mus []sync.Mutex, targetIdx, targetTerm uint64) bool {
	acks := 0
	for i, l := range logs {
		mus[i].Lock()
		if t, ok := l.TermAt(targetIdx); ok && t == targetTerm {
			acks++
		}
		mus[i].Unlock()
	}
	return acks > len(logs)/2
}
