package mod_raft_election_timer_fsm

import "sync"

type RaftNodeFSM struct {
	mu          sync.Mutex
	id          int
	peers       int
	currentTerm uint64
	votedFor    int
	role        Role
	votesGot    map[int]bool
}

func NewRaftNodeFSM(id, peers int) (*RaftNodeFSM, error) {
	if err := validateClusterSize(id, peers); err != nil {
		return nil, err
	}
	return &RaftNodeFSM{id: id, peers: peers, votedFor: -1, role: Follower, votesGot: make(map[int]bool)}, nil
}

func (r *RaftNodeFSM) StartElection() uint64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.currentTerm++
	r.role = Candidate
	r.votedFor = r.id
	r.votesGot = map[int]bool{r.id: true}
	return r.currentTerm
}

func (r *RaftNodeFSM) HandleRequestVote(candidateID int, term uint64) VoteReply {
	r.mu.Lock()
	defer r.mu.Unlock()
	if term < r.currentTerm {
		return VoteReply{Term: r.currentTerm, VoteGranted: false}
	}
	if term > r.currentTerm {
		r.currentTerm = term
		r.role = Follower
		r.votedFor = -1
	}
	if r.votedFor == -1 || r.votedFor == candidateID {
		r.votedFor = candidateID
		return VoteReply{Term: r.currentTerm, VoteGranted: true}
	}
	return VoteReply{Term: r.currentTerm, VoteGranted: false}
}

func (r *RaftNodeFSM) HandleVoteReply(from int, reply VoteReply) Role {
	r.mu.Lock()
	defer r.mu.Unlock()
	if reply.Term > r.currentTerm {
		r.currentTerm = reply.Term
		r.role = Follower
		r.votedFor = -1
		return r.role
	}
	if r.role == Candidate && reply.Term == r.currentTerm && reply.VoteGranted {
		r.votesGot[from] = true
		if len(r.votesGot) >= fastQuorum(r.peers) {
			r.role = Leader
		}
	}
	return r.role
}

func (r *RaftNodeFSM) ReceiveHeartbeat(leaderTerm uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if leaderTerm < r.currentTerm {
		return ErrRaftFsmInvalid
	}
	if leaderTerm > r.currentTerm {
		r.currentTerm = leaderTerm
		r.votedFor = -1
	}
	r.role = Follower
	return nil
}

func (r *RaftNodeFSM) Snapshot() (Role, uint64, int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.role, r.currentTerm, r.votedFor
}
