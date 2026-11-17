package mod_raft_election_timer_fsm

import "sync"

func broadcastVoteRequests(nodes []*RaftNodeFSM, candidateID int, term uint64) []VoteReply {
	replies := make([]VoteReply, len(nodes))
	var wg sync.WaitGroup
	for i, n := range nodes {
		wg.Add(1)
		go func(idx int, node *RaftNodeFSM) {
			defer wg.Done()
			replies[idx] = node.HandleRequestVote(candidateID, term)
		}(i, n)
	}
	wg.Wait()
	return replies
}
