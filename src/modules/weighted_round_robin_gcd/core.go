package mod_weighted_round_robin_gcd

type SmoothWRR struct {
	peers []PeerWeight
	total int
}

func NewSmoothWRR(weights []int) *SmoothWRR {
	w := &SmoothWRR{peers: make([]PeerWeight, len(weights))}
	for i, wt := range weights {
		w.peers[i] = PeerWeight{ID: i, Weight: wt}
		w.total += wt
	}
	return w
}

func (w *SmoothWRR) NextPeer() int {
	best := 0
	for i := range w.peers {
		w.peers[i].Current += w.peers[i].Weight
		if w.peers[i].Current > w.peers[best].Current {
			best = i
		}
	}
	w.peers[best].Current -= w.total
	return w.peers[best].ID
}
