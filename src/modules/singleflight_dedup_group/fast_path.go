package mod_singleflight_dedup_group

func (g *DedupGroup) ActiveInFlight() int {
	g.mu.Lock()
	n := len(g.calls)
	g.mu.Unlock()
	return n
}
