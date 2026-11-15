package mod_topological_task_dag

func computeCPM(durations []int64, edges []TaskEdge) *SchedulePlan {
	n := len(durations)
	adj := make([][]int, n)
	radj := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		adj[e.U] = append(adj[e.U], e.V)
		radj[e.V] = append(radj[e.V], e.U)
		indeg[e.V]++
	}
	var waves [][]int
	var topo []int
	var cur []int
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			cur = append(cur, i)
		}
	}
	for len(cur) > 0 {
		waves = append(waves, cur)
		var nxt []int
		for _, u := range cur {
			topo = append(topo, u)
			for _, v := range adj[u] {
				indeg[v]--
				if indeg[v] == 0 {
					nxt = append(nxt, v)
				}
			}
		}
		cur = nxt
	}
	if len(topo) != n {
		return &SchedulePlan{IsDAG: false}
	}
	es := make([]int64, n)
	var makespan int64
	for _, u := range topo {
		fin := es[u] + durations[u]
		if fin > makespan {
			makespan = fin
		}
		for _, v := range adj[u] {
			if fin > es[v] {
				es[v] = fin
			}
		}
	}
	ls := make([]int64, n)
	for i := range ls {
		ls[i] = makespan - durations[i]
	}
	for idx := n - 1; idx >= 0; idx-- {
		u := topo[idx]
		for _, v := range adj[u] {
			cand := ls[v] - durations[u]
			if cand < ls[u] {
				ls[u] = cand
			}
		}
	}
	_ = radj
	return &SchedulePlan{
		IsDAG:      true,
		Makespan:   makespan,
		EarlyStart: es,
		LateStart:  ls,
		Waves:      waves,
	}
}
