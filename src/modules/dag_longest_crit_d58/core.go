package mod_dag_longest_crit_d58

func DagMakespan(n int, edges []TaskEdge) (int64, bool) {
	adj := make([][]TaskEdge, n)
	ind := make([]int, n)
	for _, e := range edges {
		adj[e.U] = append(adj[e.U], e)
		ind[e.V]++
	}
	var q []int
	for i := 0; i < n; i++ {
		if ind[i] == 0 {
			q = append(q, i)
		}
	}
	dist := make([]int64, n)
	seen := 0
	for head := 0; head < len(q); head++ {
		u := q[head]
		seen++
		for _, e := range adj[u] {
			if dist[u]+e.W > dist[e.V] {
				dist[e.V] = dist[u] + e.W
			}
			ind[e.V]--
			if ind[e.V] == 0 {
				q = append(q, e.V)
			}
		}
	}
	if seen != n {
		return 0, false
	}
	var best int64
	for _, d := range dist {
		if d > best {
			best = d
		}
	}
	return best, true
}
