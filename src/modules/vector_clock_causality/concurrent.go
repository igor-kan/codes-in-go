package mod_vector_clock_causality

import "sync"

// CausalChannel models an asynchronous causal message queue guarded by Mutex.
type CausalChannel struct {
	mu   sync.Mutex
	msgs []*VectorClock
}

// Send enqueues a snapshot copy of sender's vector clock.
func (c *CausalChannel) Send(vc *VectorClock) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.msgs = append(c.msgs, vc.Clone())
}

// Drain returns all pending vector clocks.
func (c *CausalChannel) Drain() []*VectorClock {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := c.msgs
	c.msgs = nil
	return out
}
