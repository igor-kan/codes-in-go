package mod_singleflight_dedup_group

import "sync"

type flightCall struct {
	wg   sync.WaitGroup
	val  int64
	err  error
	dups int
}

type DedupGroup struct {
	mu    sync.Mutex
	calls map[string]*flightCall
}

func NewDedupGroup() *DedupGroup {
	return &DedupGroup{calls: make(map[string]*flightCall)}
}

func (g *DedupGroup) Do(key string, fn func() (int64, error)) GroupResult {
	if key == "" || fn == nil {
		return GroupResult{Err: ErrFlightEmpty}
	}
	g.mu.Lock()
	if c, ok := g.calls[key]; ok {
		c.dups++
		g.mu.Unlock()
		c.wg.Wait()
		return GroupResult{Val: c.val, Err: c.err, Shared: true}
	}
	c := new(flightCall)
	c.wg.Add(1)
	g.calls[key] = c
	g.mu.Unlock()

	val, err := fn()
	c.val = val
	c.err = err

	g.mu.Lock()
	if g.calls[key] == c {
		delete(g.calls, key)
	}
	shared := c.dups > 0
	g.mu.Unlock()
	c.wg.Done()
	return GroupResult{Val: val, Err: err, Shared: shared}
}

func (g *DedupGroup) Forget(key string) {
	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()
}
