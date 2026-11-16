package mod_reservoir_quantiles_gk

import "sync"

func BatchStreamSummaries(streams [][]float64, eps float64) ([]*GKSummary, error) {
	out := make([]*GKSummary, len(streams))
	if len(streams) == 0 {
		return out, nil
	}
	var wg sync.WaitGroup
	for i := range streams {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			g, _ := NewGKSummary(eps)
			for _, x := range streams[idx] {
				g.Observe(x)
			}
			out[idx] = g
		}(i)
	}
	wg.Wait()
	return out, nil
}
